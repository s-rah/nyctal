package wayland

import (
	"encoding/binary"
	"fmt"
)

type Display struct {
	server   *WaylandServer
	lastSync int
}

func (d *Display) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		// sync
		newID := NewUintField()
		if err := ParsePacketStructure(packet.Data, newID); err != nil {
			return err
		}
		//	d.lastSync[metadata.client] = Callback{id: uint32(*newID)}

		wsc.SendMessage(
			NewPacketBuilder(uint32(*newID), 0x00).
				WithUint(uint32(d.lastSync)).
				Build())

		wsc.SendMessage(
			NewPacketBuilder(0x01, 0x01).
				WithUint(uint32(*newID)).
				Build())

		d.lastSync += 1

		return nil
	case 1:
		// get_registry
		// 	This request creates a registry object that allows the client
		// to list and bind the global objects available from the
		//	compositor.
		newId := binary.LittleEndian.Uint32(packet.Data)
		wsc.registry.New(newId, &UnboundObject{server: d.server})

		// we only support shared memory...

		wsc.SendMessage(
			NewPacketBuilder(newId, 0x00).
				WithUint(0x01).
				WithString("wl_compositor").
				WithUint(0x05).
				Build())

		wsc.SendMessage(
			NewPacketBuilder(newId, 0x00).
				WithUint(0x02).
				WithString("wl_subcompositor").
				WithUint(0x01).
				Build())

		wsc.SendMessage(
			NewPacketBuilder(newId, 0x00).
				WithUint(0x03).
				WithString("wl_seat").
				WithUint(0x07).
				Build())

		wsc.SendMessage(
			NewPacketBuilder(newId, 0x00).
				WithUint(0x04).
				WithString("wl_shm").
				WithUint(0x02).
				Build())

		wsc.SendMessage(
			NewPacketBuilder(newId, 0x00).
				WithUint(0x05).
				WithString("xdg_wm_base").
				WithUint(0x02).
				Build())

		wsc.SendMessage(
			NewPacketBuilder(newId, 0x00).
				WithUint(0x06).
				WithString("wl_data_device_manager").
				WithUint(0x03).
				Build())

		wsc.SendMessage(
			NewPacketBuilder(newId, 0x00).
				WithUint(0x07).
				WithString("wl_output").
				WithUint(0x01).
				Build())

		// wsc.SendMessage(
		// 	NewPacketBuilder(newId, 0x00).
		// 		WithUint(0x08).
		// 		WithString("zwp_linux_dmabuf_v1").
		// 		WithUint(0x04).
		// 		Build())

		return nil
	default:
		return fmt.Errorf("unknown opcode called on display: %v", packet.Opcode)
	}

}
