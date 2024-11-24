package wayland

import (
	"fmt"

	"nyctal/model"
	"nyctal/utils"
)

type UnboundObject struct {
	BaseObject
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
			utils.Debug(int(wsc.id), "bind", fmt.Sprintf("wl_compositor#%d", new_id))
			wsc.registry.New(new_id, &Compositor{})
		case "wl_subcompositor":
			utils.Debug(int(wsc.id), "bind", fmt.Sprintf("wl_subcompositor#%d", new_id))
			wsc.registry.New(new_id, &SubCompositor{})
		case "wl_shm":
			utils.Debug(int(wsc.id), "bind", fmt.Sprintf("wl_shm#%d", new_id))
			// Send Format Message...
			wsc.SendMessage(NewPacketBuilder(new_id, 0x00).WithUint(uint32(model.FormatARGB)).Build())
			wsc.SendMessage(NewPacketBuilder(new_id, 0x00).WithUint(uint32(model.FormatXRGB)).Build())
			wsc.registry.New(new_id, &SHM{id: new_id})
		case "xdg_wm_base":
			utils.Debug(int(wsc.id), "bind", fmt.Sprintf("xdg_wm_base#%d", new_id))
			wmbase := &XDG_Base{server: u.server, wsc: wsc, id: new_id}
			wsc.pingtarget = wmbase
			wsc.registry.New(new_id, wmbase)
		case "wl_seat":
			utils.Debug(int(wsc.id), "bind", fmt.Sprintf("wl_seat#%d", new_id))
			if seat := wsc.registry.FindSeat(); seat != nil {
				seat.id = new_id
				wsc.registry.New(new_id, seat)
			} else {
				wsc.registry.New(new_id, NewSeat(wsc, new_id))
			}
		case "wl_output":
			utils.Debug(int(wsc.id), "bind", fmt.Sprintf("wl_output#%d", new_id))
			wsc.registry.New(new_id, NewOutput(new_id, wsc))
		case "wl_data_device_manager":
			utils.Debug(int(wsc.id), "bind", fmt.Sprintf("wl_data_device_manager#%d", new_id))
			wsc.registry.New(new_id, &DataDeviceManager{id: new_id})
		case "wp_viewporter":
			utils.Debug(int(wsc.id), "bind", fmt.Sprintf("wp_viewporter#%d", new_id))
			wsc.registry.New(new_id, &WPViewporter{id: new_id})
		// case "zwp_linux_dmabuf_v1":
		// 	utils.Debug("bind", fmt.Sprintf("zwp_linux_dmabuf_v1#%d", new_id))
		// 	wsc.registry.New(new_id, NewLinuxDMABuf(u.server))
		default:
			return fmt.Errorf("failed to bind: [%s]", *iname)
		}
		return nil
	default:
		return fmt.Errorf("unknown opcode called on unbound object: %v", packet.Opcode)
	}

}
