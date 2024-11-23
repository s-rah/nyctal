package wayland

import (
	"fmt"
	"image"

	"nyctal/utils"
)

type XDG_Toplevel struct {
	server *WaylandServer
	id     uint32
	title  string
	size   image.Point
}

func (u *XDG_Toplevel) Configure(wsc *WaylandServerConn, width int, height int) {

	//if width != u.size.X || height != u.size.Y {

	wsc.SendMessage(NewPacketBuilder(u.id, 0x00).
		WithUint(uint32(width)).
		WithUint(uint32(height)).
		WithUint(0).
		Build())
	u.size = image.Pt(width, height)
	//}
}

func (u *XDG_Toplevel) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		// destroy
		wsc.registry.Destroy(u.id)
		return nil
	case 1:
		// set parent
		return nil
	case 2:

		title := NewStringField()
		if err := ParsePacketStructure(packet.Data, title); err != nil {
			return err
		}

		u.title = string(*title)
		utils.Debug("xdg_toplevel", "set_title: "+u.title)
		return nil
	case 3:
		// set app_id
		return nil
	case 4:
		return nil
	case 5:
		return nil
	case 6:
		return nil // move?
	case 7:
		return nil
	case 8:
		return nil

	case 9:
		// maxiimize ignore
		return nil
	case 10:
		return nil
	case 12:
		return nil
	case 13:
		return nil
	default:
		return fmt.Errorf("unknown opcode called on xdg top level object: %v", packet.Opcode)
	}

}
