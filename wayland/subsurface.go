package wayland

import (
	"fmt"
	"image"

	"nyctal/utils"
)

type SubSurface struct {
	BaseObject
	server   *WaylandServer
	id       uint32
	surface  *Surface
	parent   *Surface
	position image.Point
	synced   bool
}

func (u *SubSurface) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		// destroy
		return nil
	case 1:
		x := NewUintField()
		y := NewUintField()
		if err := ParsePacketStructure(packet.Data, x, y); err != nil {
			return err
		}
		utils.Debug(int(wsc.id), fmt.Sprintf("subsurface#%d", u.id), fmt.Sprintf("set_position#%d", u.id))

		u.position = image.Pt(int(*x), int(*y))
		return nil
	case 2:
		// place above
		ref := NewUintField()
		if err := ParsePacketStructure(packet.Data, ref); err != nil {
			return err
		}
		utils.Debug(int(wsc.id), fmt.Sprintf("subsurface#%d", u.id), fmt.Sprintf("place_above#%d", u.id))
		return nil
	case 3:
		// place below
		ref := NewUintField()
		if err := ParsePacketStructure(packet.Data, ref); err != nil {
			return err
		}
		utils.Debug(int(wsc.id), fmt.Sprintf("subsurface#%d", u.id), fmt.Sprintf("place_below#%d", u.id))
		return nil
	case 4:
		u.synced = true
		utils.Debug(int(wsc.id), fmt.Sprintf("subsurface#%d", u.id), fmt.Sprintf("set_synced#%d", u.id))
		return nil
	case 5:
		u.synced = false
		utils.Debug(int(wsc.id), fmt.Sprintf("subsurface#%d", u.id), fmt.Sprintf("set_desynced#%d", u.id))
		return nil
	default:
		return fmt.Errorf("unknown opcode called on unbound object: %v", packet.Opcode)
	}
}
