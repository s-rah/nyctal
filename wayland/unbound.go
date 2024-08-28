package wayland

import (
	"fmt"

	"nyctal/model"
	"nyctal/utils"
)

type UnboundObject struct {
	server *WaylandServer
}

func (u *UnboundObject) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:

		inter := NewUintField()
		iname := NewStringField()
		version := NewUintField()
		newId := NewUintField()

		if err := ParsePacketStructure(packet.Data, inter, iname, version, newId); err != nil {
			return err
		}

		new_id := uint32(*newId)
		switch *iname {
		case "wl_compositor":
			utils.Debug("bind", fmt.Sprintf("wl_compositor#%d", new_id))
			wsc.registry.New(new_id, &Compositor{})
		case "wl_subcompositor":
			utils.Debug("bind", fmt.Sprintf("wl_subcompositor#%d", new_id))
			wsc.registry.New(new_id, &SubCompositor{})
		case "wl_shm":
			utils.Debug("bind", fmt.Sprintf("wl_shm#%d", new_id))
			// Send Format Message...
			wsc.SendMessage(NewPacketBuilder(new_id, 0x00).WithUint(uint32(model.FormatARGB)).Build())
			wsc.registry.New(new_id, &SHM{id: new_id})
		case "xdg_wm_base":
			utils.Debug("bind", fmt.Sprintf("xdg_wm_base#%d", new_id))
			wsc.registry.New(new_id, &XDG_Base{server: u.server, id: new_id})
		case "wl_seat":
			utils.Debug("bind", fmt.Sprintf("wl_seat#%d", new_id))
			wsc.registry.New(new_id, NewSeat(wsc, new_id))
		case "wl_output":
			utils.Debug("bind", fmt.Sprintf("wl_output#%d", new_id))
			wsc.registry.New(new_id, NewOutput(new_id, wsc))
		case "wl_data_device_manager":
			utils.Debug("bind", fmt.Sprintf("wl_data_device_manager#%d", new_id))
			wsc.registry.New(new_id, &DataDeviceManager{id: new_id})
		default:
			fmt.Printf("failed to bind: [%s]\n", *iname)
		}
		return nil
	default:
		return fmt.Errorf("unknown opcode called on unbound object: %v", packet.Opcode)
	}

}
