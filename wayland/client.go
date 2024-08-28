package wayland

import (
	"fmt"
	"image"

	"nyctal/model"
	"nyctal/utils"
)

type WaylandClient struct {
	id           int
	surface      *XDG_Surface
	hasPointer   bool
	popups       utils.Stack[*XDGPopup]
	pointerLocal image.Point
	wsc          *WaylandServerConn
}

func NewWaylandClient(client int, wsc *WaylandServerConn, surface *XDG_Surface) model.Client {
	return &WaylandClient{wsc: wsc, id: client, surface: surface}
}

func (wc *WaylandClient) PushPopup(popup *XDGPopup) {
	wc.popups.Push(popup)
}

func (wc *WaylandClient) PopPopup() *XDGPopup {
	top, _ := wc.popups.Pop()
	return top
}

func (wc *WaylandClient) GetChild(id int) model.Client {
	return nil
}

func (wc *WaylandClient) Resize(width int, height int) {
	utils.Debug(fmt.Sprintf("wayland_client#%d", wc.id), fmt.Sprintf("resizing %d %d %v", width, height, wc.surface))
	if wc.surface.windowGeometry.Dx() != width || wc.surface.windowGeometry.Dy() != height {
		wc.surface.Resize(wc.wsc, width, height)
	}
}

func (wc *WaylandClient) AckFrame() {
	//wc.surface.AckFrame(wc.id)
}

func (wc *WaylandClient) Parent() model.Client {
	return nil
}

func (wc *WaylandClient) AddChild(id int, child model.Client) int {
	// todo use this interface?
	panic("unimplemented add child for window")
}

func (wc *WaylandClient) RemoveChild(id int) {

}

func (wc *WaylandClient) RemoveChildren(id int) {

}

func (wc *WaylandClient) Buffer() *model.BGRA {
	wl_surface := wc.surface.surface
	img := wl_surface.cached
	if img != nil {
		wg := wc.surface.windowGeometry
		if wg.Dx() == 0 {
			wg = img.Bounds()
		}

		atZero := image.Rect(0, 0, wg.Dx(), wg.Dy())
		window := model.EmptyBGRA(atZero)
		model.DrawCopyOver(window, atZero, img, wg.Min)
		wc.surface.surface.RenderFrame(wc.wsc)

		for _, client := range wc.popups.Inner() {
			if !client.configured {
				client.Configure(wc.wsc)
				continue
			}

			xdg_surface := client.surface

			wl_surface := xdg_surface.surface

			pimg := wl_surface.cached
			if pimg != nil {
				wl_surface.RenderFrame(wc.wsc)
				offset := xdg_surface.RelativeOffset()
				atZero := image.Rect(offset.X, offset.Y,
					offset.X+client.positioner.size.Dx(),
					offset.Y+client.positioner.size.Dy())
				//utils.Debug("client", fmt.Sprintf("rendering at %v %v\n", atZero, pimg.Bounds()))
				model.DrawCopyOver(window, atZero, pimg, image.Pt(xdg_surface.windowGeometry.Min.X, xdg_surface.windowGeometry.Min.Y))
			} else {
				utils.Debug("client", "could not render popup...")
			}

		}

		if wc.hasPointer {
			if seat := wc.wsc.registry.FindSeat(); seat != nil {
				if pointer, err := wc.wsc.registry.Get(seat.mouse); err == nil {
					if pointerObj, ok := pointer.(*Pointer); ok {
						//utils.Debug("client", fmt.Sprintf("drawing cursor %v", pointerObj.local))
						ps, _ := wc.wsc.registry.Get(pointerObj.surface)
						if pointer_surface, ok := ps.(*Surface); ok {
							mouseBuf := pointer_surface.cached
							if mouseBuf != nil {
								pointerImgLoc := wc.pointerLocal.Sub(pointerObj.hotspot)
								windowRect := image.Rect(pointerImgLoc.X, pointerImgLoc.Y, pointerImgLoc.X+mouseBuf.Bounds().Dx(), pointerImgLoc.Y+mouseBuf.Bounds().Dy())
								model.DrawCopyOver(window, windowRect, mouseBuf, image.Pt(0, 0))
								pointer_surface.RenderFrame(wc.wsc)
							}
						}
					}
				}
			}
		}

		return window
	} else {
		utils.Debug("client", "parent surface buffer was nil...")
	}
	return nil
}

func (wc *WaylandClient) ProcessKeyboardEvent(ev model.KeyboardEvent) {
	seat := wc.wsc.registry.FindSeat()
	if seat != nil {
		seat.ProcessKeyboardEvent(ev)
	}
}

func (wc *WaylandClient) ProcessPointerEvent(ev model.PointerEvent) bool {

	// send pointer enter event
	seat := wc.wsc.registry.FindSeat()
	if seat != nil {
		if ev.Move != nil {
			wc.pointerLocal = image.Pt(int(ev.Move.MX), int(ev.Move.MY))
		}
		seat.ProcessPointerEvent(wc.wsc, ev, wc.surface)
		wc.hasPointer = true
		return true
	}
	return false
}

func (wc *WaylandClient) ProcessFocus() {
	seat := wc.wsc.registry.FindSeat()
	if seat != nil {
		seat.Grab(wc.surface)
	}
}

func (wc *WaylandClient) ProcessUnFocus() {

}

func (wc *WaylandClient) HandlePointerLeave() {
	if wc.hasPointer {
		// send pointer leave evner
		wc.hasPointer = false
	}
}
