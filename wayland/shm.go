package wayland

import (
	"fmt"

	"nyctal/utils"
)

type SHM struct {
	server *WaylandServer
	id     uint32
}

func (u *SHM) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		newId := NewUintField()
		size := NewUintField()
		if err := ParsePacketStructure(packet.Data, newId, size); err != nil {
			return err
		}

		utils.Debug(fmt.Sprintf("shm#%d", u.id), fmt.Sprintf("wm_shm_pool#%d %d", *newId, *size))

		fd, ok := wsc.fds.Pop()
		fmt.Printf("pop fd (%d)\n", fd)
		if !ok {
			return fmt.Errorf("no fd in queue")
		}

		pool, err := NewSHMPool(uint32(*newId), u.server, fd, uint32(*size))
		if err != nil {
			return err
		}
		wsc.registry.New(uint32(*newId), pool)
		return nil

	default:
		return fmt.Errorf("unknown opcode called on unbound object: %v", packet.Opcode)
	}

}
