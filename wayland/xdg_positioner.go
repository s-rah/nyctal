package wayland

import (
	"fmt"
	"image"

	"nyctal/utils"
)

type XDG_Positioner struct {
	BaseObject
	id                   uint32
	size                 image.Rectangle
	anchorRect           image.Rectangle
	anchor               uint32
	gravity              uint32
	constraintAdjustment uint32
	offset               image.Point
	reactive             bool
}

func (u *XDG_Positioner) CalculateAnchorPoint() image.Point {
	if u.anchor == 5 { // top left
		return u.offset.Add(image.Pt(u.anchorRect.Min.X, u.anchorRect.Min.Y))
	} else if u.anchor == 6 { // bottom left
		return u.offset.Add(image.Pt(u.anchorRect.Min.X, u.anchorRect.Max.Y))
	} else if u.anchor == 7 { // top right
		return u.offset.Add(image.Pt(u.anchorRect.Max.X, u.anchorRect.Min.Y))
	} else if u.anchor == 8 { // bottom rights
		return u.offset.Add(image.Pt(u.anchorRect.Max.X, u.anchorRect.Max.Y))
	} else {
		return u.offset.Add(image.Pt(u.anchorRect.Min.X+u.anchorRect.Dx()/2, u.anchorRect.Min.Y+u.anchorRect.Dy()/2))
	}
}

func (u *XDG_Positioner) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		// destroy
		wsc.registry.Destroy(u.id)
		return nil
	case 1:
		w := NewUintField()
		h := NewUintField()
		if err := ParsePacketStructure(packet.Data, w, h); err != nil {
			return err
		}
		utils.Debug(int(wsc.id), fmt.Sprintf("xdg_positioner#%d", u.id), fmt.Sprintf("set_size %d %d", uint32(*w), uint32(*h)))
		u.size = image.Rect(0, 0, int(*w), int(*h))
		return nil
	case 2:
		x := NewIntField()
		y := NewIntField()
		w := NewIntField()
		h := NewIntField()
		if err := ParsePacketStructure(packet.Data, x, y, w, h); err != nil {
			return err
		}
		utils.Debug(int(wsc.id), fmt.Sprintf("xdg_positioner#%d", u.id), fmt.Sprintf("set_anchor_rect %d %d %d %d", int32(*x), int32(*y), int32(*w), int32(*h)))
		u.anchorRect = image.Rect(int(*x), int(*y), int(*x)+int(*w), int(*y)+int(*h))
		return nil
	case 3:
		x := NewUintField()
		if err := ParsePacketStructure(packet.Data, x); err != nil {
			return err
		}
		utils.Debug(int(wsc.id), fmt.Sprintf("xdg_positioner#%d", u.id), fmt.Sprintf("set_anchor %d", uint32(*x)))
		u.anchor = uint32(*x)
		return nil
	case 4:
		x := NewUintField()
		if err := ParsePacketStructure(packet.Data, x); err != nil {
			return err
		}
		utils.Debug(int(wsc.id), fmt.Sprintf("xdg_positioner#%d", u.id), fmt.Sprintf("set_gravity %d", uint32(*x)))
		u.gravity = uint32(*x)
		return nil
	case 5:
		x := NewUintField()
		if err := ParsePacketStructure(packet.Data, x); err != nil {
			return err
		}
		utils.Debug(int(wsc.id), fmt.Sprintf("xdg_positioner#%d", u.id), fmt.Sprintf("set_constraint_adjustment %d", uint32(*x)))
		u.constraintAdjustment = uint32(*x)
		return nil
	case 6:
		x := NewIntField()
		y := NewIntField()
		if err := ParsePacketStructure(packet.Data, x, y); err != nil {
			return err
		}
		utils.Debug(int(wsc.id), fmt.Sprintf("xdg_positioner#%d", u.id), fmt.Sprintf("offset %d %d", int32(*x), int32(*y)))
		u.offset = image.Pt(int(*x), int(*y))
		return nil
	case 7:
		utils.Debug(int(wsc.id), fmt.Sprintf("xdg_positioner#%d", u.id), "set_reactive")
		u.reactive = true
		return nil
	case 8:

		return nil
	case 9:

		return nil
	default:
		return fmt.Errorf("unknown opcode called on xdg_positioner: %v", packet.Opcode)
	}

}
