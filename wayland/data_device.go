package wayland

import "fmt"

type DataDevice struct {
	BaseObject
	server    *WaylandServer
	id        uint32
	seat      *Seat
	selection *DataSource
}

func (u *DataDevice) Selection(wsc *WaylandServerConn) {
	wsc.SendMessage(NewPacketBuilder(u.id, 0x05).WithUint(0).Build())
}

func (u *DataDevice) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 1:
		// 	This request asks the compositor to set the selection
		//to the data from the source on behalf of the client.

		//To unset the selection, set the source to NULL.

		//The given source may not be used in any further set_selection or
		//start_drag requests. Attempting to reuse a previously-used source
		//may send a used_source error.
		source := NewUintField()
		serial := NewUintField()
		if err := ParsePacketStructure(packet.Data, source, serial); err != nil {
			return err
		}
		if obj, err := wsc.registry.Get((uint32(*source))); err == nil {
			if datasource, ok := obj.(*DataSource); ok {
				u.selection = datasource
				return nil
			}
		}
		return fmt.Errorf("could not set selection")
	default:
		return fmt.Errorf("unknown opcode called on data device: %v", packet.Opcode)
	}

}
