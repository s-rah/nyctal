package wayland

import (
	"crypto/rand"
	"fmt"
	"image"

	"nyctal/model"
	"nyctal/utils"
)

type Surface struct {
	id            uint32
	frameCallback uint32

	attached bool
	pending  *Buffer

	commitedInputRegion *Region
	pendingInputRegion  *Region

	parent   uint32
	children []uint32
	cached   *model.BGRA
	damage   []image.Rectangle
	first    bool
}

func (u *Surface) AddChild(child_surface uint32) {
	u.children = append(u.children, child_surface)
}

func (u *Surface) read_buffer() *model.BGRA {

	if u.pending != nil && u.pending.format == model.FormatARGB {

		wl_pool := u.pending.backing
		if wl_pool != nil {

			bounds := image.Rect(0, 0, int(u.pending.width), int(u.pending.height))
			if u.cached == nil || len(u.damage) == 0 || bounds != u.cached.Bounds() {

				img := model.NewBGRA(wl_pool.data[u.pending.offset:], bounds, int(u.pending.stride))
				u.cached = img
			} else {

				for _, damage := range u.damage {
					if damage.Min.Y < 0 {
						damage.Min.Y = 0
					}

					if damage.Max.Y > int(u.pending.height) {
						damage.Max.Y = int(u.pending.height)
					}

					if damage.Min.X < 0 {
						damage.Min.X = 0
					}

					if damage.Max.X > int(u.pending.width) {
						damage.Max.X = int(u.pending.width)
					}
					u.cached.Update(wl_pool.data[u.pending.offset:], damage, int(u.pending.stride))
				}
			}
			return u.cached
		}
		utils.Debug("surface", "unable to find shmpool")
	}
	utils.Debug("surface", "unable to init buffer")
	return nil
}

func (u *Surface) RenderFrame(wsc *WaylandServerConn) {

	if u.frameCallback != 0 {
		serial := []byte{0, 0, 0, 0}
		rand.Read(serial)
		wsc.SendMessage(
			NewPacketBuilder(u.frameCallback, 0x00).
				WithBytes(serial).
				Build())

		utils.Debug(fmt.Sprintf("xdg_surface#%d", u.id), fmt.Sprintf("callback frame#%d", u.frameCallback))

		wsc.SendMessage(
			NewPacketBuilder(0x01, 0x01).
				WithUint(u.frameCallback).
				Build())

	}

	u.frameCallback = 0
}

func (u *Surface) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		// destroy
		wsc.registry.Destroy(u.id)
		return nil
	case 1:
		// Set a buffer as the content of this surface.
		// Surface contents are double-buffered state, see wl_surface.commit.
		bufferId := NewUintField()
		x := NewIntField()
		y := NewIntField()
		if err := ParsePacketStructure(packet.Data, bufferId, x, y); err != nil {
			return err
		}

		if uint32(*bufferId) == 0 {
			// If wl_surface.attach is sent with a NULL wl_buffer, the
			// following wl_surface.commit will remove the surface content.
			u.pending = nil
			u.attached = true
			return nil
		}

		utils.Debug(fmt.Sprintf("surface#%d", u.id), fmt.Sprintf("attach_buffer#%d %d %d", *bufferId, *x, *y))
		if obj, err := wsc.registry.Get(uint32(*bufferId)); err == nil {
			if buffer, ok := obj.(*Buffer); ok {
				u.pending = buffer
				u.attached = true
				return nil
			}
		}

		return fmt.Errorf("failed to attach buffer: unknown buffer")
	case 2:
		x := NewIntField()
		y := NewIntField()
		w := NewIntField()
		h := NewIntField()
		if err := ParsePacketStructure(packet.Data, x, y, w, h); err != nil {
			return err
		}
		//fmt.Printf("set window geometry?\n")
		utils.Debug(fmt.Sprintf("surface#%d", u.id), fmt.Sprintf("damage %d %d %d %d", *x, *y, *w, *h))
		u.damage = append(u.damage, image.Rect(int(*x), int(*y), int(*x)+int(*w), int(*y)+int(*h)))
		return nil
	case 3:
		newId := NewUintField()
		if err := ParsePacketStructure(packet.Data, newId); err != nil {
			return err
		}
		utils.Debug(fmt.Sprintf("surface#%d", u.id), fmt.Sprintf("frame_callback#%d", *newId))
		u.frameCallback = uint32(*newId)
		return nil
	case 4:
		//fmt.Printf("set input region\n")
		regionId := NewUintField()
		if err := ParsePacketStructure(packet.Data, regionId); err != nil {
			return err
		}
		utils.Debug(fmt.Sprintf("surface#%d", u.id), fmt.Sprintf("set_input_region#%d", *regionId))
		rid := uint32(*regionId)
		if rid == 0 {
			u.pendingInputRegion = nil
			return nil
		} else {
			if obj, err := wsc.registry.Get(rid); err == nil {
				if region, ok := obj.(*Region); ok {
					u.pendingInputRegion = region
					return nil
				}
			}
		}
		return fmt.Errorf("unknown region reference")
	case 5:
		// we ignore all opauqe region hints...
		return nil
	case 8:
		//fmt.Printf("set buffer scale\n")
		return nil
	case 6:

		u.commitedInputRegion = u.pendingInputRegion

		// received create pool message...
		if u.pending != nil {

			// Committing a pending wl_buffer allows the compositor to read the
			// pixels in the wl_buffer. The compositor may access the pixels at
			// any time after the wl_surface.commit request. When the compositor
			// will not access the pixels anymore, it will send the
			// wl_buffer.release event. Only after receiving wl_buffer.release,
			// the client may reuse the wl_buffer. A wl_buffer that has been
			// attached and then replaced by another attach instead of committed
			// will not receive a release event, and is not used by the
			// compositor.
			u.read_buffer()

			u.pending.Release(wsc)

		} else if u.attached {
			// If wl_surface.attach is sent with a NULL wl_buffer, the
			// following wl_surface.commit will remove the surface content.
			u.cached = nil
		}

		// After commit, there is no pending buffer until the next attach.
		u.attached = false
		u.pending = nil

		utils.Debug(fmt.Sprintf("surface#%d", u.id), "commit")

		// Pretend we are entering a surface....
		if !u.first {
			u.first = true
			output := wsc.registry.FindOutput()
			wsc.SendMessage(
				NewPacketBuilder(u.id, 0x00).WithUint(output.id).
					Build())
		}

		return nil
	case 9:
		return nil
	case 10:
		return nil
	default:
		return fmt.Errorf("unknown opcode called on surface: %v", packet.Opcode)
	}
}
