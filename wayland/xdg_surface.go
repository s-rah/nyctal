package wayland

import (
	"fmt"
	"image"

	"nyctal/model"
	"nyctal/utils"
)

type XDG_Surface struct {
	id   uint32          // client given id
	uniq model.GlobalIdx // compositor unique id

	server         *WaylandServer
	surface        *Surface
	topLevel       *XDG_Toplevel
	positioner     *XDG_Positioner
	windowGeometry image.Rectangle
	hasPointer     bool

	parent *XDG_Surface
	popup  *XDGPopup
	offset image.Point

	configuring bool
	serial      uint32
}

func (u *XDG_Surface) RelativeOffset() image.Point {
	pOffset := image.Pt(0, 0)
	if u.parent != nil {
		pOffset = u.parent.offset
	}
	return pOffset.Add(u.offset).Sub(u.windowGeometry.Min)
}

// takes in top-level surface-local coordinates and checks if the pointer
// intersects the popup...
func (xp *XDG_Surface) Intersects(pointer image.Point) bool {
	if xp.surface == nil {
		return false // un undefed surface cannot be intersected
	}
	if xp.positioner == nil {
		if xp.surface.commitedInputRegion == nil {
			return true
		} else {
			inArea, inRegion := xp.surface.commitedInputRegion.Intersects(pointer)
			if inArea && inRegion {
				return true
			} else if !inArea {
				return true
			}
		}
		return false
	}

	tl := xp.RelativeOffset()
	size := xp.positioner.size
	bounds := image.Rect(tl.X, tl.Y, tl.X+size.Dx(), tl.Y+size.Dy())

	if pointer.In(bounds) {
		// we are inside this popup
		surfaceLocal := pointer.Add(tl)
		if xp.surface.commitedInputRegion == nil {
			return true
		} else {
			inArea, inRegion := xp.surface.commitedInputRegion.Intersects(surfaceLocal)
			if inArea && inRegion {
				return true
			} else if !inArea {
				return true
			}
		}
	}
	return false
}

func (u *XDG_Surface) Configure(wsc *WaylandServerConn) {

	if !u.configuring {

		if u.topLevel != nil {
			wsc.SendMessage(
				NewPacketBuilder(u.topLevel.id, 0x03).
					WithUint(0).
					Build())

		}

		u.serial += 1
		wsc.SendMessage(
			NewPacketBuilder(u.id, 0x00).
				WithUint(u.serial).
				Build())

		utils.Debug(fmt.Sprintf("xdg_surface#%d", u.id), "configure")
		u.configuring = true
	}
}

func (u *XDG_Surface) Resize(wsc *WaylandServerConn, width int, height int) {

	if u.topLevel != nil {
		u.topLevel.Configure(wsc, width, height)
	}
	u.Configure(wsc)

}

func (u *XDG_Surface) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		wsc.registry.Destroy(u.id)
		u.server.workspace.RemoveTopLevel(u.uniq)
		return nil
	case 1:
		new_id := NewUintField()
		if err := ParsePacketStructure(packet.Data, new_id); err != nil {
			return err
		}
		utils.Debug(fmt.Sprintf("xdg_surface#%d", u.id), fmt.Sprintf("get_xdg_toplevel#%d", uint32(*new_id)))
		topLevel := &XDG_Toplevel{server: u.server, id: uint32(*new_id)}
		wsc.registry.New(uint32(*new_id), topLevel)
		u.topLevel = topLevel

		u.uniq = model.GlobalIdx(wsc.registry.globalIdx.Add(1))
		u.server.workspace.AddTopLevel(NewWaylandClient(u.uniq, wsc.id, wsc, u))

		if seat := wsc.registry.FindSeat(); seat != nil {
			seat.Grab(u)
			return nil
		} else {
			return fmt.Errorf("top_level: could not find seat")
		}
	case 2:
		new_id := NewUintField()
		surface := NewUintField()
		positioner := NewUintField()
		if err := ParsePacketStructure(packet.Data, new_id, surface, positioner); err != nil {
			return err
		}
		utils.Debug(fmt.Sprintf("xdg_surface#%d", u.id), fmt.Sprintf("get_popup#%d %d %d", uint32(*new_id), *surface, &positioner))

		if surfaceObj, err := wsc.registry.Get(uint32(*surface)); err != nil {
			return err
		} else if parentSurface, ok := surfaceObj.(*XDG_Surface); ok {

			wl_positioner, _ := wsc.registry.Get(uint32(*positioner))
			if positioner, ok := wl_positioner.(*XDG_Positioner); ok {
				utils.Debug(fmt.Sprintf("xdg_surface#%d", u.id), fmt.Sprintf("setting_offset#%v", positioner.anchorRect))
				u.offset = positioner.CalculateAnchorPoint()

				popup := &XDGPopup{server: u.server, id: uint32(*new_id), parent: parentSurface, surface: u, positioner: positioner}
				u.parent = parentSurface
				u.positioner = positioner
				u.parent.popup = popup
				wsc.registry.New(uint32(*new_id), popup)

				// find the parent window...
				window := u.server.workspace.GetTopLevel(parentSurface.uniq)
				// setup popup rendering in the client
				if wayland, ok := window.(*WaylandClient); ok {
					wayland.PushPopup(popup)
				}

				// send a configure event to the client
				popup.Configure(wsc)
				return nil
			}
		}
		return fmt.Errorf("could not setup xdgpopup")
	case 3:
		x := NewIntField()
		y := NewIntField()
		w := NewUintField()
		h := NewUintField()
		if err := ParsePacketStructure(packet.Data, x, y, w, h); err != nil {
			return err
		}
		utils.Debug(fmt.Sprintf("xdg_surface#%d", u.id), fmt.Sprintf("set_window_geometry %d %d %d %d", uint32(*x), uint32(*y), uint32(*w), uint32(*h)))
		u.windowGeometry = image.Rect(int(*x), int(*y), int(*x)+int(*w), int(*y)+int(*h))
		return nil
	case 4:
		// ack confgure
		u.configuring = false
		return nil
	default:
		return fmt.Errorf("unknown opcode called on xdg surface object: %v", packet.Data)
	}

}
