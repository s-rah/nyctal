package wayland

import (
	"fmt"
	"nyctal/model"
)

type Buffer struct {
	BaseObject
	id          uint32
	wsc         *WaylandServerConn
	backingPool *SHMPool
	width       uint32
	height      uint32
	stride      uint32
	offset      uint32
	format      model.Format
	destroyed   bool
}

func (u *Buffer) Destroy() {
	if !u.destroyed {
		// u.wsc.SendMessage(
		// 	NewPacketBuilder(u.id, 0x00).
		// 		Build())
		// u.wsc.SendMessage(
		// 	NewPacketBuilder(0x1, 0x01).WithUint(u.id).
		// 		Build())
		u.destroyed = true
	}

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
