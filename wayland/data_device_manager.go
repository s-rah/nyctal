package wayland

import (
	"fmt"

	"nyctal/utils"
)

type DataDeviceManager struct {
	BaseObject
	server *WaylandServer
	id     uint32
}

func (u *DataDeviceManager) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		newId := NewUintField()
		if err := ParsePacketStructure(packet.Data, newId); err != nil {
			return err
		}

		utils.Debug(int(wsc.id), "data_device_manager", fmt.Sprintf("create_data_source#%d", *newId))
		wsc.registry.New(uint32(*newId), &DataSource{id: uint32(*newId), mimetypes: make(map[string]bool)})
		return nil
	case 1:

		newId := NewUintField()
		seatId := NewUintField()
		if err := ParsePacketStructure(packet.Data, newId, seatId); err != nil {
			return err
		}

		utils.Debug(int(wsc.id), "data_device_manager", fmt.Sprintf("get_data_device#%d %d", *newId, *seatId))

		if obj, err := wsc.registry.Get((uint32(*seatId))); err == nil {
			if seat, ok := obj.(*Seat); ok {
				wsc.registry.New(uint32(*newId), &DataDevice{id: uint32(*newId), seat: seat, server: u.server})
				return nil
			}
		}

		return fmt.Errorf("unable to get data device: invalid seat")
	default:
		return fmt.Errorf("unknown opcode called on data device manager: %v", packet.Opcode)
	}

}
