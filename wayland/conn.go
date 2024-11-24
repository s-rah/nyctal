package wayland

import (
	"encoding/binary"
	"fmt"
	"net"
	"sync/atomic"

	"syscall"

	"golang.org/x/sys/unix"

	"nyctal/model"
	"nyctal/utils"
)

type Pingable interface {
	Ping()
}

type WaylandServerConn struct {
	socket     net.Conn
	registry   *Registry
	id         model.GlobalIdx
	fds        *utils.Queue[int]
	connFd     int
	index      *atomic.Uint32
	pingtarget Pingable
	errors     int
}

func (c *WaylandServerConn) SendMessageWithFd(data []byte, fd int) {
	utils.Debug(int(c.id), "send-wayland-message-with-fd", fmt.Sprintf("%d %x", c.id, data))
	rights := syscall.UnixRights([]int{fd}...)
	syscall.Sendmsg(c.connFd, data, rights, nil, 0)
}

func (c *WaylandServerConn) SendMessage(data []byte) {
	utils.Debug(int(c.id), "send-wayland-message", fmt.Sprintf("%d %x", c.id, data))
	syscall.Sendmsg(c.connFd, data, nil, nil, 0)
}

func (c *WaylandServerConn) RecvMsg(connFd int, p []byte) (int, error) {

	b := make([]byte, unix.CmsgSpace(4))

	syscall.SetsockoptTimeval(connFd, syscall.SOL_SOCKET, syscall.SO_RCVTIMEO, &syscall.Timeval{Sec: 2})
	n, oobn, flags, _, err := syscall.Recvmsg(connFd, p, b, 0)

	// parse socket control message
	if oobn > 0 {
		cmsgs, err := unix.ParseSocketControlMessage(b)
		if err != nil || cmsgs == nil {
			return -1, fmt.Errorf("ERROR CMSG %v %v %v", cmsgs, flags, err)
		}
		fds, err := unix.ParseUnixRights(&cmsgs[0])
		if err != nil {
			return -1, fmt.Errorf("ERROR PUR %v %v %v", cmsgs, flags, err)
		}
		for _, fd := range fds {
			c.fds.Push(fd)
		}

	}

	return n, err
}

func (c *WaylandServerConn) ReadPacket() (*WaylandMessage, error) {

	p := make([]byte, 8)

	n, err := c.RecvMsg(c.connFd, p)

	if err != nil {
		return nil, fmt.Errorf("expected 8 got %d: %v", n, err)
	}

	if n == -1 {
		return nil, fmt.Errorf("returned -1")
	}

	if n == 0 {
		return nil, fmt.Errorf("returned 0")
	}

	if n != 8 {
		return nil, fmt.Errorf("unexpectd number of bytes")
	}

	msg := &WaylandMessage{}
	msg.Address = binary.LittleEndian.Uint32(p)
	msg.Length = binary.LittleEndian.Uint16(p[6:8])
	msg.Opcode = binary.LittleEndian.Uint16(p[4:6])
	msg.Length -= 8
	if msg.Length > 0 {
		p = make([]byte, msg.Length)
		n, err = c.RecvMsg(c.connFd, p)
		if err != nil {
			return nil, err
		}

		if n == -1 {
			return nil, fmt.Errorf("none")
		}

		if n != int(msg.Length) {
			return nil, fmt.Errorf("expected %d got %d", n, msg.Length)
		}
		msg.Data = p
	}
	return msg, nil
}
