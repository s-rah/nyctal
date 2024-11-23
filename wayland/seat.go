package wayland

import (
	"fmt"
	"image"

	"nyctal/model"
	"nyctal/utils"
)

type Seat struct {
	server       *WaylandServer
	id           uint32
	keyboard     *Keyboard
	mouse        *Pointer
	serial       uint32
	DataDevice   *DataDevice
	pointerFocus *XDG_Surface
}

func (s *Seat) Grab(surface *XDG_Surface) {
	if s.keyboard != nil {
		s.keyboard.Leave(s.serial)
		s.serial += 1
		s.keyboard.Enter(s.serial, surface.surface)
	}
}

func (s *Seat) ProcessKeyboardEvent(ev model.KeyboardEvent) {
	if s.keyboard != nil {
		s.serial += 1
		s.keyboard.ProcessKeyboardEvent(ev, s.serial)
	}
}

func checkIntersect(point image.Point, surface *XDG_Surface) *XDG_Surface {

	if surface == nil {
		return nil
	}

	var highestIntersect *XDG_Surface
	if surface.popup != nil {
		highestIntersect = checkIntersect(point, surface.popup.surface)
	}

	if highestIntersect == nil {
		if surface.Intersects(point) {
			return surface
		}
	}
	return highestIntersect
}

func (s *Seat) ProcessPointerEvent(wsc *WaylandServerConn, ev model.PointerEvent, surface *XDG_Surface) {

	if s.mouse != nil && ev.Move != nil {

		is := checkIntersect(image.Pt(int(ev.Move.MX), int(ev.Move.MY)), surface)

		// we intersected with nothing
		if is == nil {
			if s.pointerFocus != nil {
				s.serial += 1
				utils.Debug(fmt.Sprintf("wl_pointer#%d", s.mouse), fmt.Sprintf("leave %d ", s.pointerFocus.surface.id))
				wsc.SendMessage(NewPacketBuilder(s.mouse.id, 0x01).
					WithUint(s.serial).
					WithUint(s.pointerFocus.surface.id).Build())
				s.pointerFocus.hasPointer = false
				s.pointerFocus = nil
			}
			return
		}

		if s.pointerFocus == nil {
			s.pointerFocus = is
			s.pointerFocus.hasPointer = false
		} else if is.id != s.pointerFocus.id {
			s.serial += 1
			utils.Debug(fmt.Sprintf("wl_pointer#%d", s.mouse), fmt.Sprintf("leave %d ", s.pointerFocus.surface.id))
			wsc.SendMessage(NewPacketBuilder(s.mouse.id, 0x01).
				WithUint(s.serial).
				WithUint(s.pointerFocus.surface.id).Build())
			utils.Debug(fmt.Sprintf("wl_pointer#%d", s.mouse), fmt.Sprintf("leave %d ", s.pointerFocus.surface.id))
			s.pointerFocus.hasPointer = false
			s.pointerFocus = is
		}
	}
	top := s.pointerFocus
	if top != nil {

		if ev.Move != nil {
			// these coordinates are in top_level local coords, we need to
			// convert them to surface level
			ev.Move.MX -= float32(top.RelativeOffset().X)
			ev.Move.MY -= float32(top.RelativeOffset().Y)
		}

		if s.mouse != nil {

			if !top.hasPointer {

				if s.keyboard != nil {
					s.serial += 1
					s.keyboard.Leave(s.serial)
					s.serial += 1
					s.keyboard.Enter(s.serial, top.surface)
				}

				s.serial += 1
				pb := NewPacketBuilder(s.mouse.id, 0x00).
					WithUint(s.serial).
					WithUint(top.surface.id)

				if ev.Move != nil {
					pb = pb.WithFixed(ev.Move.MX)
					pb = pb.WithFixed(ev.Move.MY)
				} else {
					pb = pb.WithFixed(0)
					pb = pb.WithFixed(0)

				}
				wsc.SendMessage(pb.Build())
				utils.Debug(fmt.Sprintf("wl_pointer#%d", s.mouse), fmt.Sprintf("enter %d %v", top.surface.id, ev.Move))
				top.hasPointer = true
				wsc.SendMessage(NewPacketBuilder(s.mouse.id, 0x05).Build())

			}

			if ev.Move != nil {
				wsc.SendMessage(NewPacketBuilder(s.mouse.id, 0x02).
					WithUint(ev.Move.Time).
					WithFixed(ev.Move.MX).
					WithFixed(ev.Move.MY).Build())
				//utils.Debug(fmt.Sprintf("wl_pointer#%d", s.mouse), "motion")
			}

			if ev.Button != nil {
				s.serial += 1
				wsc.SendMessage(NewPacketBuilder(s.mouse.id, 0x03).
					WithUint(s.serial).
					WithUint(ev.Button.Time).
					WithUint(ev.Button.Button).
					WithUint(ev.Button.State).Build())
				//utils.Debug(fmt.Sprintf("wl_pointer#%d", s.mouse), "button")
			}

			if ev.Axis != nil {
				wsc.SendMessage(NewPacketBuilder(s.mouse.id, 0x04).
					WithUint(ev.Axis.Time).
					WithUint(ev.Axis.Axis).
					WithFixed(ev.Axis.Value).Build())
				//utils.Debug(fmt.Sprintf("wl_pointer#%d", s.mouse), "axis")
			}
			wsc.SendMessage(NewPacketBuilder(s.mouse.id, 0x05).Build())
			//	utils.Debug(fmt.Sprintf("wl_pointer#%d", s.mouse), "frame")
		}
	}
}

func NewSeat(wsc *WaylandServerConn, id uint32) *Seat {

	wsc.SendMessage(NewPacketBuilder(id, 0x00).WithUint(0x03).Build())
	utils.Debug(fmt.Sprintf("wl_seat#%d", id), "capabilities")
	wsc.SendMessage(NewPacketBuilder(id, 0x01).WithString("default").Build())

	return &Seat{id: id}
}

func (u *Seat) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:

		mouse_id := NewUintField()
		if err := ParsePacketStructure(packet.Data, mouse_id); err != nil {
			return err
		}
		u.mouse = &Pointer{server: u.server, id: uint32(*mouse_id)}
		fmt.Printf("newseatmouse id %v\n", u.mouse)

		wsc.registry.New(uint32(*mouse_id), u.mouse)

		utils.Debug("wl_seat", fmt.Sprintf("get_pointer#%d", u.mouse))
		return nil
	case 1:
		keyboard_id := NewUintField()
		if err := ParsePacketStructure(packet.Data, keyboard_id); err != nil {
			return err
		}
		kbid := uint32(*keyboard_id)
		u.keyboard = NewKeyboard(kbid, wsc)
		utils.Debug("wl_seat", fmt.Sprintf("get_keyboard#%d", u.keyboard.id))
		return nil
	default:
		return fmt.Errorf("unknown opcode called on wl_seat: %v", packet.Opcode)
	}

}
