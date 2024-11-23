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

		if surface, err := wsc.registry.Get(uint32(*surface)); err == nil {
			if surfaceObj, ok := surface.(*Surface); ok {

				if parentsurface, err := wsc.registry.Get(uint32(*parent)); err == nil {
					if parentsurfaceObj, ok := parentsurface.(*Surface); ok {

						subSurface := &SubSurface{server: u.server, id: uint32(*newId), surface: surfaceObj, parent: parentsurfaceObj}
						wsc.registry.New(uint32(*newId), subSurface)
						parentsurfaceObj.AddSubSurface(subSurface)
						return nil
					}
				}
			}
		}

		return fmt.Errorf("could not set up subsurface")
	default:
		return fmt.Errorf("unknown opcode called on unbound object: %v", packet.Opcode)
	}
}
