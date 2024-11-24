package wayland

import (
	"fmt"
)

type Output struct {
	BaseObject
	id uint32
}

func NewOutput(id uint32, wsc *WaylandServerConn) *Output {
	output := &Output{id: id}
	wsc.registry.New(id, output)

	wsc.SendMessage(NewPacketBuilder(id, 0x00).
		WithUint(0).
		WithUint(0).
		WithUint(1024).
		WithUint(1024).
		WithUint(0).
		WithString("nyctal").
		WithString("none").
		WithUint(0).
		Build())

	wsc.SendMessage(NewPacketBuilder(id, 0x01).
		WithUint(3).WithUint(1024).WithUint(1024).WithUint(60000).Build())
	wsc.SendMessage(NewPacketBuilder(id, 0x03).WithUint(1).Build())
	wsc.SendMessage(NewPacketBuilder(id, 0x02).Build())
	return output
}

func (u *Output) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		// destroy
		return nil
	case 1:
		return nil
	default:
		return fmt.Errorf("unknown opcode called on region: %v", packet.Opcode)
	}

}
