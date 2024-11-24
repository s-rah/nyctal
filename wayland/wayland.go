package wayland

import (
	"fmt"
	"net"
	"nyctal/model"
	"nyctal/utils"
	"runtime/debug"
	"strings"
	"sync/atomic"
	"syscall"

	"golang.org/x/sys/unix"
)

// WaylandServer encapsulates everything to do with the wayland protocol
// On the network side clients connect to the server and send events
// On the service side the user (via the OS) send input updates
type WaylandServer struct {
	l         net.Listener
	socket    string
	globalIdx atomic.Uint32
	workspace model.Workspace
}

func NewServer(display_socket string, workspace model.Workspace) (*WaylandServer, error) {
	l, err := net.Listen("unix", display_socket)
	if err != nil {
		return nil, err
	}
	ws := &WaylandServer{socket: display_socket,
		l:         l,
		workspace: workspace,
	}

	return ws, nil
}

func (ws *WaylandServer) Listen() {
	clientId := 0
	for {
		fd, err := ws.l.Accept()
		clientId += 1
		if err != nil {
			println("accept error", err.Error())
			return
		}
		connFd, err := getConnFd(fd.(*net.UnixConn))
		if err != nil {
			return
		}

		wsc := &WaylandServerConn{
			socket:   fd,
			index:    &ws.globalIdx,
			connFd:   connFd,
			id:       model.GlobalIdx(clientId),
			fds:      utils.NewQueue[int](),
			registry: NewRegistry(),
		}
		go ws.handle(wsc)
	}
}

func (ws *WaylandServer) handle(wsc *WaylandServerConn) {

	// Ok..we need to do this because its the only way of
	// robustly handling malicious/broken clients who truncate
	// or otherwise break the shm contract.
	// Instead of crashing the whole application, go will instead
	// panic this goroutine, which we can then safely recover from
	// (see below...)
	debug.SetPanicOnFault(true)

	wsc.registry.New(0, &NullObject{})
	wsc.registry.New(1, &Display{lastSync: 5, server: ws})
	utils.Debug(int(wsc.id), "ws", fmt.Sprintf("new client#%d", wsc.id))

	defer func() {

		if r := recover(); r != nil {
			debug.PrintStack()
			utils.Debug(int(wsc.id), "wayland-server", fmt.Sprintf("recovered panic in client thread %v", r))
		}
		utils.Debug(int(wsc.id), "wayland-server", fmt.Sprintf("client#%d removed", wsc.id))
		ws.workspace.RemoveAllWithParent(wsc.id)
		wsc.registry.Close()

		wsc.socket.Close()
		for !wsc.fds.Empty() {
			fd, _ := wsc.fds.Pop()
			unix.Close(fd)
		}

	}()

	for {

		packet, err := wsc.ReadPacket()
		if err != nil {
			utils.Debug(int(wsc.id), "wayland-server", err.Error())
			if strings.Contains(err.Error(), "resource temporarily unavailable") && wsc.errors < 10 {
				if wsc.pingtarget != nil {
					wsc.pingtarget.Ping()
					wsc.errors += 1
					continue
				}
			}
			break
		} else {
			utils.Debug(int(wsc.id), "wayland-message", fmt.Sprintf("%d %v", wsc.id, packet))
		}

		if obj, err := wsc.registry.Get(uint32(packet.Address)); err == nil {
			if err := obj.HandleMessage(wsc, packet); err != nil {
				utils.Debug(int(wsc.id), "client", err.Error())
				break
			}
		} else {
			utils.Debug(int(wsc.id), "client", err.Error())
			break
		}

	}
	utils.Debug(int(wsc.id), "client", "terminating")

}

func getConnFd(conn syscall.Conn) (connFd int, err error) {
	var rawConn syscall.RawConn
	rawConn, err = conn.SyscallConn()
	if rawConn == nil || err != nil {
		return
	}

	err = rawConn.Control(func(fd uintptr) {
		connFd = int(fd)
	})
	syscall.SetNonblock(connFd, false)
	return
}
