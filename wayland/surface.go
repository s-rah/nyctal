package wayland

import (
	"fmt"
	"image"
	"time"

	"nyctal/model"
	"nyctal/utils"
)

type Surface struct {
	BaseObject
	id            uint32
	frameCallback utils.Queue[uint32]

	attached bool
	pending  *Buffer

	commitedInputRegion *Region
	pendingInputRegion  *Region

	children []*SubSurface
	cached   *model.BGRA
	damage   []image.Rectangle
	first    bool
}

func (u *Surface) AddSubSurface(child_surface *SubSurface) {
	u.children = append(u.children, child_surface)
}

func (u *Surface) Destroy() {
	u.cached = nil
}

func (u *Surface) RenderBuffer() {
	// Committing a pending wl_buffer allows the compositor to read the
	// pixels in the wl_buffer. The compositor may access the pixels at
	// any time after the wl_surface.commit request. When the compositor
	// will not access the pixels anymore, it will send the
	// wl_buffer.release event. Only after receiving wl_buffer.release,
	// the client may reuse the wl_buffer. A wl_buffer that has been
	// attached and then replaced by another attach instead of committed
	// will not receive a release event, and is not used by the
	// compositor.
	if u.pending != nil {
		u.read_buffer()
		u.pending.Destroy()
		u.pending = nil
	}
}

func (u *Surface) read_buffer() *model.BGRA {

	if u.pending != nil && u.pending.format == model.FormatARGB {

		wl_pool := u.pending.backingPool
		if wl_pool != nil {

			bounds := image.Rect(0, 0, int(u.pending.width), int(u.pending.height))
			if u.cached == nil || len(u.damage) == 0 || bounds != u.cached.Bounds() {
				if bounds.Dy()*int(u.pending.stride) > 2048*2048*16 {
					return nil
				}

				data := wl_pool.mappedData
				img := model.NewBGRA(data[u.pending.offset:], bounds, int(u.pending.stride))
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
					u.cached.Update((wl_pool.mappedData)[u.pending.offset:], damage, int(u.pending.stride))
				}
			}
			return u.cached
		}
	}
	return nil
}

func (u *Surface) RenderFrame(wsc *WaylandServerConn, serial []byte) []byte {

	for !u.frameCallback.Empty() {
		nullserial := uint32(time.Now().UnixMilli())

		cb, _ := u.frameCallback.Pop()
		wsc.SendMessage(
			NewPacketBuilder(cb, 0x00).
				WithUint(nullserial).
				Build())

		utils.Debug(int(wsc.id), fmt.Sprintf("xdg_surface#%d", u.id), fmt.Sprintf("callback frame#%d", u.frameCallback))

		wsc.SendMessage(
			NewPacketBuilder(0x01, 0x01).
				WithUint(cb).
				Build())

	}

	return serial
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

		utils.Debug(int(wsc.id), fmt.Sprintf("surface#%d", u.id), fmt.Sprintf("attach_buffer#%d %d %d", *bufferId, *x, *y))
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
		utils.Debug(int(wsc.id), fmt.Sprintf("surface#%d", u.id), fmt.Sprintf("damage %d %d %d %d", *x, *y, *w, *h))
		u.damage = append(u.damage, image.Rect(int(*x), int(*y), int(*x)+int(*w), int(*y)+int(*h)))
		return nil
	case 3:
		newId := NewUintField()
		if err := ParsePacketStructure(packet.Data, newId); err != nil {
			return err
		}
		utils.Debug(int(wsc.id), fmt.Sprintf("surface#%d", u.id), fmt.Sprintf("frame_callback#%d", *newId))
		u.frameCallback.Push(uint32(*newId))
		return nil
	case 4:
		regionId := NewUintField()
		if err := ParsePacketStructure(packet.Data, regionId); err != nil {
			return err
		}
		utils.Debug(int(wsc.id), fmt.Sprintf("surface#%d", u.id), fmt.Sprintf("set_input_region#%d", *regionId))
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
	case 7:
		// set buffer something...
		return nil
	case 8:
		return nil
	case 6:
		u.commitedInputRegion = u.pendingInputRegion

		// received create pool message...
		if u.pending != nil {

		} else if u.attached {
			// If wl_surface.attach is sent with a NULL wl_buffer, the
			// following wl_surface.commit will remove the surface content.
			u.cached = nil
		}

		// After commit, there is no pending buffer until the next attach.
		u.attached = false

		utils.Debug(int(wsc.id), fmt.Sprintf("surface#%d", u.id), "commit")

		// // Pretend we are entering a surface....
		// if !u.first {
		// 	u.first = true
		// 	output := wsc.registry.FindOutput()
		// 	if output != nil && !u.first {
		// 		u.first = true
		// 		wsc.SendMessage(
		// 			NewPacketBuilder(u.id, 0x00).WithUint(output.id).
		// 				Build())
		// 	}
		// 	dd := wsc.registry.FindDataDevice()
		// 	if dd != nil {
		// 		dd.Selection(wsc)
		// 	}
		// }

		return nil
	case 9:
		return nil
	case 10:
		return nil
	default:
		return fmt.Errorf("unknown opcode called on surface: %v", packet.Opcode)
	}
}
