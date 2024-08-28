package wayland

import (
	"fmt"


	"nyctal/utils"
)

type SubCompositor struct {
	server *WaylandServer
}

func (u *SubCompositor) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		// destroy
		return nil
	case 1:
		newId := NewUintField()
		surface := NewUintField()
		parent := NewUintField()
		if err := ParsePacketStructure(packet.Data, newId, surface, parent); err != nil {
			return err
		}
		utils.Debug("compositor", fmt.Sprintf("create_subsurface#%d", newId))

		wsc.registry.New(uint32(*newId),
			&SubSurface{server: u.server, id: uint32(*newId), surface: uint32(*surface), parent: uint32(*parent)})

		if surface, err := wsc.registry.Get(uint32(*surface)); err == nil {
			if surfaceObj, ok := surface.(*Surface); ok {
				surfaceObj.parent = uint32(*parent)
				surfaceObj.AddChild(uint32(*newId))
			} else {
				return fmt.Errorf("could not find linekd surface")
			}
		} else {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unknown opcode called on unbound object: %v", packet.Opcode)
	}
}
