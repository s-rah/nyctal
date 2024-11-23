package wayland

import (
	"fmt"
	"net"
	"nyctal/model"
	"nyctal/utils"
	"runtime/debug"
	"syscall"
)

// WaylandServer encapsulates everything to do with the wayland protocol
// On the network side clients connect to the server and send events
// On the service side the user (via the OS) send input updates
type WaylandServer struct {
	l      net.Listener
	socket string

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

		wsc := &WaylandServerConn{
			socket:   fd,
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
	utils.Debug("ws", fmt.Sprintf("new client#%d", wsc.id))

	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			fmt.Println("Recovered panic in client thread:", r)
		}
		utils.Debug("wayland-server", fmt.Sprintf("client#%d removed", wsc.id))
		ws.workspace.RemoveAllWithParent(wsc.id)
	}()

	for {

		packet, err := wsc.ReadPacket()
		if err != nil {
			utils.Debug("wayland-server", err.Error())
			break
		} else {
			utils.Debug("wayland-message", fmt.Sprintf("%d %v", wsc.id, packet))
		}

		if obj, err := wsc.registry.Get(uint32(packet.Address)); err == nil {
			if err := obj.HandleMessage(wsc, packet); err != nil {
				fmt.Printf("[error] %v\n", err)
				break
			}
		} else {
			fmt.Printf("[error] %v\n", err)
			break
		}

	}

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
	return
}
