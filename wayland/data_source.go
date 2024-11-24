package wayland

import "fmt"

type DataSource struct {
	BaseObject
	id        uint32
	mimetypes map[string]bool
}

func (u *DataSource) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		// 	This request adds a mime type to the set of mime types
		// advertised to targets.  Can be called several times to offer
		// multiple types.
		mimetype := NewStringField()
		if err := ParsePacketStructure(packet.Data, mimetype); err != nil {
			return err
		}
		u.mimetypes[string(*mimetype)] = true
		return nil
	case 1:
		// 	Destroy the data source.
		wsc.registry.Destroy(u.id)
		return nil
	default:
		return fmt.Errorf("unknown opcode called on data source: %v", packet.Opcode)
	}

}
