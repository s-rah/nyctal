package wayland

import (
	"fmt"

	"golang.org/x/sys/unix"

	"nyctal/utils"
)

type SHMPool struct {
	id     uint32
	server *WaylandServer
	fd     int
	size   uint32
	data   []byte
}

func NewSHMPool(id uint32, server *WaylandServer, fd int, size uint32) (*SHMPool, error) {
	data, err := unix.Mmap(fd, 0, int(size), unix.PROT_READ, unix.MAP_SHARED)
	if err != nil {
		return nil, err
	}
	return &SHMPool{id: id, server: server, fd: fd, size: size, data: data}, nil
}

func (u *SHMPool) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:

		newId := NewUintField()
		offset := NewUintField()
		width := NewUintField()
		height := NewUintField()
		stride := NewUintField()
		if err := ParsePacketStructure(packet.Data, newId, offset, width, height, stride); err != nil {
			return err
		}

		utils.Debug("shm_pool", fmt.Sprintf("create_buffer#%d %d %d %d %d ", *newId, *offset, *width, *height, *stride))
		wsc.registry.New(uint32(*newId),
			&Buffer{id: uint32(*newId), server: u.server, backing: u, offset: uint32(*offset), stride: uint32(*stride), width: uint32(*width), height: uint32(*height)})
		return nil
	case 1:
		// destroy
		wsc.registry.Destroy(u.id)
		// wsc.SendMessage(
		// 	NewPacketBuilder(0x1, 0x01).WithUint(u.id).
		// 		Build())

		return nil
	case 2:
		// resize
		newsize := NewUintField()
		if err := ParsePacketStructure(packet.Data, newsize); err != nil {
			return err
		}
		u.size = uint32(*newsize)
		utils.Debug("shm_pool", fmt.Sprintf("resize, %d ", *newsize))

		data, err := unix.Mremap(u.data, int(u.size), unix.MREMAP_MAYMOVE)
		if err != nil {
			return fmt.Errorf("could not remap data: %v", err)
		}
		u.data = data

		return nil

	default:
		return fmt.Errorf("unknown opcode called on shmpool: %v", packet.Opcode)
	}
}
