package wayland

import "fmt"

type WPViewporter struct {
	id uint32
}

func (u *WPViewporter) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {
	return fmt.Errorf("unknown opcode called on viewporter object: %v", packet.Opcode)
}
