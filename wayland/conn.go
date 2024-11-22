package wayland

import (
	"encoding/binary"
	"fmt"
	"net"

	"syscall"
	"time"

	"golang.org/x/sys/unix"

	"nyctal/model"
	"nyctal/utils"
)

type WaylandServerConn struct {
	socket   net.Conn
	registry *Registry
	id       model.GlobalIdx
	fds      *utils.Queue[int]
}

func (c *WaylandServerConn) SendMessageWithFd(data []byte, fd int) {
	connFd, _ := getConnFd(c.socket.(*net.UnixConn))
	rights := syscall.UnixRights([]int{fd}...)
	syscall.Sendmsg(connFd, data, rights, nil, 0)
}

func (c *WaylandServerConn) SendMessage(data []byte) {
	c.socket.Write(data)
}

func (c *WaylandServerConn) RecvMsg(connFd int, p []byte) (int, error) {

	b := make([]byte, unix.CmsgSpace(4))
	syscall.SetNonblock(connFd, false)
	n, oobn, _, _, err := syscall.Recvmsg(connFd, p, b, 0)

	// parse socket control message
	if oobn > 0 {
		cmsgs, err := unix.ParseSocketControlMessage(b)
		if err != nil || cmsgs == nil {
			return 0, err
		}
		fds, err := unix.ParseUnixRights(&cmsgs[0])
		if err != nil {
			return 0, err
		}
		for _, fd := range fds {
			fmt.Printf("pushing fd (%d)\n", fd)
			c.fds.Push(fd)
		}
	}

	if n == -1 {
		time.Sleep(time.Millisecond)
		return c.RecvMsg(connFd, p)
	}
	return n, err
}

func (c *WaylandServerConn) ReadPacket() (*WaylandMessage, error) {

	connFd, err := getConnFd(c.socket.(*net.UnixConn))
	if err != nil {
		return nil, fmt.Errorf("could not get conn fd: %v", err)
	}

	p := make([]byte, 8)

	n, err := c.RecvMsg(connFd, p)

	if n == -1 {
		return c.ReadPacket()
	}

	if n == 0 {
		return nil, fmt.Errorf("nothing")
	}
	if err != nil || n != 8 {
		return nil, fmt.Errorf("expected 8 got %d: %v", n, err)
	}
	msg := &WaylandMessage{}
	msg.Address = binary.LittleEndian.Uint32(p)
	msg.Length = binary.LittleEndian.Uint16(p[6:8])
	msg.Opcode = binary.LittleEndian.Uint16(p[4:6])
	msg.Length -= 8
	if msg.Length > 0 {
		p = make([]byte, msg.Length)
		n, err = c.RecvMsg(connFd, p)
		if err != nil {
			return nil, err
		}

		if n == -1 {
			return c.ReadPacket()
		}

		if n != int(msg.Length) {
			return nil, fmt.Errorf("expected %d got %d", n, msg.Length)
		}
		msg.Data = p
	}
	return msg, nil
}
