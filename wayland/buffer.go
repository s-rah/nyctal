package wayland

import (
	"fmt"
	"nyctal/model"
)

type Buffer struct {
	id      uint32
	server  *WaylandServer
	backing *SHMPool
	width   uint32
	height  uint32
	stride  uint32
	offset  uint32
	format  model.Format
}

func (u *Buffer) Release(wsc *WaylandServerConn) {

	// wsc.SendMessage(
	// 	NewPacketBuilder(u.id, 0x00).
	// 		Build())
	// wsc.SendMessage(
	// 	NewPacketBuilder(0x1, 0x01).WithUint(u.id).
	// 		Build())
}

func (u *Buffer) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		wsc.registry.Destroy(u.id)
		return nil
	default:
		return fmt.Errorf("unknown opcode called on buffer: %v", packet.Opcode)
	}
}
