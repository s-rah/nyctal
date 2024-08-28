package wayland

import (
	"fmt"

	"nyctal/utils"
)

type Compositor struct {
}

func (u *Compositor) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		newId := NewUintField()
		if err := ParsePacketStructure(packet.Data, newId); err != nil {
			return err
		}
		utils.Debug("compositor", fmt.Sprintf("create_surface#%d", *newId))
		wsc.registry.New(uint32(*newId), &Surface{id: uint32(*newId)})
		return nil
	case 1:
		newId := NewUintField()
		if err := ParsePacketStructure(packet.Data, newId); err != nil {
			return err
		}
		wsc.registry.New(uint32(*newId), NewRegion(uint32(*newId)))
		return nil
	default:
		return fmt.Errorf("unknown opcode called on unbound object: %v", packet.Opcode)
	}
}
