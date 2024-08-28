package wayland

import (
	"fmt"
	"image"

	"nyctal/utils"
)

type Pointer struct {
	server  *WaylandServer
	id      uint32
	surface uint32
	hotspot image.Point // cached pointer hotspot coords..
}

func (u *Pointer) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		// set_cursor
		serial := NewUintField()
		surface := NewUintField()
		x := NewIntField()
		y := NewIntField()
		//The parameters hotspot_x and hotspot_y define the position of
		//the pointer surface relative to the pointer location. Its
		//top-left corner is always at (x, y) - (hotspot_x, hotspot_y),
		//where (x, y) are the coordinates of the pointer location, in
		//surface-local coordinates.
		if err := ParsePacketStructure(packet.Data, serial, surface, x, y); err != nil {
			return err
		}
		utils.Debug(fmt.Sprintf("pointer#%d", u.id), fmt.Sprintf("set_cursor %d %d %d %d", *serial, *surface, *x, *y))
		u.surface = uint32(*surface)
		u.hotspot = image.Pt(int(*x), int(*y))

		return nil
	case 1:
		return nil
	default:
		return fmt.Errorf("unknown opcode called on pointer: %v", packet.Opcode)
	}

}
