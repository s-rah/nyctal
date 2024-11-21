package wayland

import (
	"fmt"
	"nyctal/utils"
)

// NOTE: Prototype Code...
type LinuxDMABuf struct {
	server *WaylandServer
}

func NewLinuxDMABuf(server *WaylandServer) *LinuxDMABuf {
	return &LinuxDMABuf{server: server}
}

func (u *LinuxDMABuf) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		// destroy
		return nil
	case 1:
		newId := NewUintField()
		if err := ParsePacketStructure(packet.Data, newId); err != nil {
			return err
		}
		wsc.registry.New(uint32(*newId), NewLinuxDMABufParams(uint32(*newId), wsc))
		return nil
	case 2:
		newId := NewUintField()
		if err := ParsePacketStructure(packet.Data, newId); err != nil {
			return err
		}
		wsc.registry.New(uint32(*newId), NewLinuxDMABufFeedback(uint32(*newId), 0, wsc))
		return nil
	case 3:
		newId := NewUintField()
		surfaceId := NewUintField()
		if err := ParsePacketStructure(packet.Data, newId, surfaceId); err != nil {
			return err
		}
		wsc.registry.New(uint32(*newId), NewLinuxDMABufFeedback(uint32(*newId), uint32(*surfaceId), wsc))
		return nil
	default:
		return fmt.Errorf("unknown opcode called on linux dmabuf object: %v", packet.Opcode)
	}
}

type LinuxDMABufParams struct {
	id uint32
}

func NewLinuxDMABufParams(id uint32, wsc *WaylandServerConn) *LinuxDMABufParams {
	return &LinuxDMABufParams{}
}

func (u *LinuxDMABufParams) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		// destroy
		return nil
	default:
		return fmt.Errorf("unknown opcode called on unbound object: %v", packet.Opcode)
	}
}

type LinuxDMABufFeedback struct {
	id        uint32
	surfaceId uint32
}

func NewLinuxDMABufFeedback(id uint32, surfaceId uint32, wsc *WaylandServerConn) *LinuxDMABufFeedback {
	formatTable := []byte{'X', 'R', '2', '4', 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		'A', 'R', '2', '4', 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	formatFile, _, _ := utils.Memfile("format", formatTable)

	wsc.SendMessageWithFd(NewPacketBuilder(id, 0x01).WithUint(uint32(len(formatTable))).Build(), int(formatFile)) // send format table..
	wsc.SendMessage(NewPacketBuilder(id, 0x02).WithU64Array([]uint64{0xE200}).Build())                            // send main device
	wsc.SendMessage(NewPacketBuilder(id, 0x04).WithU64Array([]uint64{0xE200}).Build())                            // send tranche table..
	wsc.SendMessage(NewPacketBuilder(id, 0x05).WithU16Array([]uint16{0x00, 0x01}).Build())
	wsc.SendMessage(NewPacketBuilder(id, 0x06).WithUint(0).Build())
	wsc.SendMessage(NewPacketBuilder(id, 0x03).Build()) // tranche done
	wsc.SendMessage(NewPacketBuilder(id, 0x00).Build()) // tranche done
	return &LinuxDMABufFeedback{id: id, surfaceId: surfaceId}
}

func (u *LinuxDMABufFeedback) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		// destroy
		return nil

	default:
		return fmt.Errorf("unknown opcode called on linux dmabuf feedback object: %v", packet.Opcode)
	}
}
