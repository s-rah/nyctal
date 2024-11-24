package wayland

import (
	"fmt"
	"time"

	"nyctal/utils"
)

type XDG_Base struct {
	BaseObject
	server *WaylandServer
	wsc    *WaylandServerConn
	id     uint32
}

func (u *XDG_Base) Ping() {
	u.wsc.SendMessage(NewPacketBuilder(u.id, 0x00).WithUint(uint32(time.Now().UnixMilli())).Build())
}

func (u *XDG_Base) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		wsc.registry.Destroy(u.id)
		return nil
	case 1:
		//create positioner
		new_id := NewUintField()
		if err := ParsePacketStructure(packet.Data, new_id); err != nil {
			return err
		}
		utils.Debug(int(wsc.id), "xdg_wm_base", fmt.Sprintf("create_positioner %d", *new_id))
		xdg_positioner := &XDG_Positioner{id: uint32(*new_id)}
		wsc.registry.New(uint32(*new_id), xdg_positioner)
		return nil
	case 2:
		new_id := NewUintField()
		surface_id := NewUintField()
		if err := ParsePacketStructure(packet.Data, new_id, surface_id); err != nil {
			return err
		}

		// received create pool message...
		utils.Debug(int(wsc.id), "xdg_wm_base", fmt.Sprintf("get_xdg_surface %d %d", new_id, surface_id))

		if surface, err := wsc.registry.Get(uint32(*surface_id)); err == nil {
			if surfaceObj, ok := surface.(*Surface); ok {
				xdgsurface := &XDG_Surface{server: u.server, surface: surfaceObj, id: uint32(*new_id)}
				wsc.registry.New(uint32(*new_id), xdgsurface)
			} else {
				return fmt.Errorf("object is not a surface")
			}
		} else {
			return err
		}

		return nil

	default:
		return fmt.Errorf("unknown opcode called on xdg: %v", packet.Opcode)
	}

}
