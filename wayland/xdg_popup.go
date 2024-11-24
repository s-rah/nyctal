package wayland

import (
	"fmt"

	"nyctal/utils"
)

type XDGPopup struct {
	BaseObject
	id         uint32
	server     *WaylandServer
	parent     *XDG_Surface
	surface    *XDG_Surface
	positioner *XDG_Positioner

	configured bool
}

func (xp *XDGPopup) Configure(wsc *WaylandServerConn) {

	wsc.SendMessage(NewPacketBuilder(xp.id, 0x00).
		WithUint(uint32(xp.positioner.anchorRect.Min.X)).
		WithUint(uint32(xp.positioner.anchorRect.Min.Y)).
		WithUint(uint32(xp.positioner.size.Dx())).
		WithUint(uint32(xp.positioner.size.Dy())).
		Build())
	xp.configured = true
	utils.Debug(int(wsc.id), fmt.Sprintf("xdg_popup#%d", xp.id), "configure")

	xp.surface.Configure(wsc)
	xp.parent.Configure(wsc)

}

func (u *XDGPopup) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		// destroy
		wsc.registry.Destroy(u.id)
		window := u.server.workspace.GetTopLevel(u.parent.uniq)
		if wayland, ok := window.(*WaylandClient); ok {
			wayland.PopPopup()
			u.parent.popup = nil // prevent new surface intersections TODO: improve this interface
		}
		return nil
	case 1:
		x := NewUintField()
		if err := ParsePacketStructure(packet.Data, x); err != nil {
			return err
		}
		utils.Debug(int(wsc.id), fmt.Sprintf("xdg_popup#%d", u.id), fmt.Sprintf("grab seat#%d", uint32(*x)))
		gseat := uint32(*x)
		if seat := wsc.registry.FindSeat(); seat != nil && seat.id == gseat {
			seat.Grab(u.surface)
			return nil
		} else {
			return fmt.Errorf("grab_seat: could not find seat %v", seat)
		}
	case 2:
		// reposition
		u.parent.Configure(wsc)
		return nil
	default:
		return fmt.Errorf("unknown opcode called on xdg_popup: %v", packet.Opcode)
	}

}
