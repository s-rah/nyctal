package wayland

import (
	"fmt"
	"image"
	"nyctal/utils"
	"slices"
)

type Area struct {
	rect     image.Rectangle
	subtract bool
}

//	A region object describes an area.
//
// Region objects are used to describe the opaque and input  regions of a surface.
type Region struct {
	id    uint32
	rects *utils.Stack[Area]
}

func NewRegion(id uint32) *Region {
	return &Region{id: id, rects: utils.NewStack[Area]()}
}

// Returns two booleans. The first indicates if the point intersects with a  sub-area
// the second indicates if the area was part of the defined region.
// This is needed because input areas and opaque areas are defined by different defaults.
func (u *Region) Intersects(point image.Point) (bool, bool) {
	utils.Debug(fmt.Sprintf("region#%d", u.id), fmt.Sprintf("computing intersection %v", point))
	areas := u.rects.Inner()
	slices.Reverse(areas)
	for _, area := range areas {
		if point.In(area.rect) {
			return true, !area.subtract
		}
	}

	return false, false
}

func (u *Region) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		wsc.registry.Destroy(u.id)
		// destroy
		wsc.SendMessage(
			NewPacketBuilder(0x1, 0x01).WithUint(u.id).
				Build())
		return nil
	case 1:
		// Add the specified rectangle to the region.
		x := NewIntField()
		y := NewIntField()
		width := NewIntField()
		height := NewIntField()
		if err := ParsePacketStructure(packet.Data, x, y, width, height); err != nil {
			return err
		}
		u.rects.Push(Area{rect: image.Rect(int(*x), int(*y), int(*x)+int(*width), int(*y)+int(*height)), subtract: false})
		return nil
	case 2:
		// Substract the specified rectangle to the region.
		x := NewIntField()
		y := NewIntField()
		width := NewIntField()
		height := NewIntField()
		if err := ParsePacketStructure(packet.Data, x, y, width, height); err != nil {
			return err
		}
		u.rects.Push(Area{rect: image.Rect(int(*x), int(*y), int(*x)+int(*width), int(*y)+int(*height)), subtract: true})
		return nil
	default:
		return fmt.Errorf("unknown opcode called on region: %v", packet.Opcode)
	}

}
