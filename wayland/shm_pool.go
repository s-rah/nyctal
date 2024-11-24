package wayland

import (
	"fmt"

	"nyctal/utils"

	"golang.org/x/sys/unix"
)

type SHMPool struct {
	id  uint32
	wsc *WaylandServerConn

	size       uint32
	mappedData []byte
}

func NewSHMPool(id uint32, wsc *WaylandServerConn, fd int, size uint32) (*SHMPool, error) {
	defer unix.Close(fd)
	if size > 2048*2048*16 {
		return nil, fmt.Errorf("attempting to make pool of huge size: %v", size/1024/1024)
	}
	data, err := unix.Mmap(fd, 0, int(size), unix.PROT_READ, unix.MAP_SHARED)
	if err != nil {
		return nil, err
	}

	pool := &SHMPool{id: id, wsc: wsc, size: size, mappedData: data}
	utils.Debug(int(wsc.id), "shmpool", fmt.Sprintf("created pool %v %p", pool.id, pool.mappedData))
	return pool, nil
}

func (u *SHMPool) Destroy() {

	if u.mappedData != nil {
		utils.Debug(int(u.wsc.id), "shmpool", fmt.Sprintf("destroying pool %v %p", u.id, u.mappedData))
		err := unix.Munmap(u.mappedData)
		u.mappedData = nil
		utils.Debug(int(u.wsc.id), "shmpool", fmt.Sprintf("destroying pool %v %v", u.id, err))
	}
}

func (u *SHMPool) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:

		if u.mappedData == nil {
			return fmt.Errorf("pool has not been initialized")
		}

		newId := NewUintField()
		offset := NewUintField()
		width := NewUintField()
		height := NewUintField()
		stride := NewUintField()
		if err := ParsePacketStructure(packet.Data, newId, offset, width, height, stride); err != nil {
			return err
		}

		utils.Debug(int(wsc.id), "shm_pool", fmt.Sprintf("create_buffer#%d %d %d %d %d ", *newId, *offset, *width, *height, *stride))
		wsc.registry.New(uint32(*newId),
			&Buffer{id: uint32(*newId), wsc: wsc, backingPool: u, offset: uint32(*offset), stride: uint32(*stride), width: uint32(*width), height: uint32(*height)})
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
		utils.Debug(int(wsc.id), "shm_pool", fmt.Sprintf("resize, %d ", *newsize))

		data, err := unix.Mremap(u.mappedData, int(u.size), unix.MREMAP_MAYMOVE)
		if data == nil || err != nil {
			return fmt.Errorf("could not remap data: %v", err)
		}
		u.mappedData = data

		return nil

	default:
		return fmt.Errorf("unknown opcode called on shmpool: %v", packet.Opcode)
	}
}
