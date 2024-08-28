package wayland

import (
	"fmt"
)

type Region struct {
	width  int
	height int
	id     uint32
}

func (u *Region) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		wsc.registry.Destroy(u.id)
		// destroy
		wsc.SendMessage(
			NewPacketBuilder(0x1, 0x01).WithUint(u.id).
				Build())
		return nil
	case 1:
		return nil
	default:
		return fmt.Errorf("unknown opcode called on region: %v", packet.Opcode)
	}

}
