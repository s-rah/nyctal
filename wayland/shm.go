package wayland

import (
	"fmt"

	"nyctal/utils"
)

type SHM struct {
	BaseObject
	id uint32
}

func (u *SHM) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		newId := NewUintField()
		size := NewUintField()
		if err := ParsePacketStructure(packet.Data, newId, size); err != nil {
			return err
		}

		utils.Debug(int(wsc.id), fmt.Sprintf("shm#%d", u.id), fmt.Sprintf("wm_shm_pool#%d %d", *newId, *size))

		fd, err := wsc.fds.Pop()
		if err != nil {
			return fmt.Errorf("expected an fd, but could not pop from queue: %v", err)
		}

		pool, err := NewSHMPool(uint32(*newId), wsc, fd, uint32(*size))
		if err != nil {
			return err
		}
		wsc.registry.New(uint32(*newId), pool)
		return nil

	default:
		return fmt.Errorf("unknown opcode called on unbound object: %v", packet.Opcode)
	}

}
