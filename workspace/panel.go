package workspace

import (
	"fmt"
	"image"
	"nyctal/model"
	"nyctal/utils"
	"runtime/debug"
	"sync"
)

type Panel struct {
	windows *utils.Queue[model.TopLevelWindow]
	bounds  image.Rectangle
	lock    sync.Mutex
	do      *DragOverlay
}

func NewWindowPanel(do *DragOverlay) model.Workspace {
	return &Panel{do: do, windows: utils.NewQueue[model.TopLevelWindow]()}
}

func (p *Panel) AddTopLevel(window model.TopLevelWindow) {
	p.lock.Lock()
	defer p.lock.Unlock()
	if window == nil {
		debug.PrintStack()
		panic("")
	}
	utils.Debug("panel", fmt.Sprintf("adding top window: %v", window))
	p.windows.PushTop(window)
}

func (p *Panel) GetTopLevel(idx model.GlobalIdx) model.TopLevelWindow {
	p.lock.Lock()
	defer p.lock.Unlock()
	for _, window := range p.windows.Inner() {
		if window.Index() == idx {
			return window
		}
	}
	return nil
}

func (p *Panel) RemoveTopLevel(idx model.GlobalIdx) {
	p.lock.Lock()
	defer p.lock.Unlock()
	utils.Debug("panel", fmt.Sprintf("remove top window: %v", idx))
	newWindows := utils.NewQueue[model.TopLevelWindow]()
	for !p.windows.Empty() {
		window, _ := p.windows.Pop()
		if window.Index() != idx {
			newWindows.Push(window)
		}
	}
	p.windows = newWindows
}

func (p *Panel) RemoveAllWithParent(pidx model.GlobalIdx) {
	p.lock.Lock()
	defer p.lock.Unlock()
	utils.Debug("panel", fmt.Sprintf("remove parent window: %v", pidx))
	newWindows := utils.NewQueue[model.TopLevelWindow]()
	for !p.windows.Empty() {
		window, _ := p.windows.Pop()
		if window.Parent() != pidx {
			newWindows.Push(window)
		}
	}
	p.windows = newWindows
}

func (p *Panel) Buffer(img *model.BGRA, width int, height int) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if top, exists := p.windows.Top(); exists {
		top.Buffer(img, width, height)
	}

}

func (p *Panel) ProcessKeyboardEvent(pointer model.Pointer, kb model.Keyboard, ev model.KeyboardEvent) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if top, exists := p.windows.Top(); exists {
		top.ProcessKeyboardEvent(ev)
	}
}

func (p *Panel) ProcessPointerEvent(pointer model.Pointer, kb model.Keyboard, ev model.PointerEvent) bool {
	p.lock.Lock()
	defer p.lock.Unlock()

	if !p.windows.Empty() {
		downKeys := kb.DownKeys()
		if downKeys[model.KB_CTRL] && downKeys[model.KB_ALT] {
			if ev.Button != nil && p.do.dragging == nil {
				if ev.Button.Button == 0x110 && ev.Button.State == 0x01 {
					window, _ := p.windows.Pop()
					p.do.Start(pointer, window)
					return true
				}
			}
		}
	}

	if top, exists := p.windows.Top(); exists {
		if ev.Move != nil {
			ev.Move.MX = float32(pointer.MX)
			ev.Move.MY = float32(pointer.MY)
		}
		return top.ProcessPointerEvent(ev)
	}
	return false
}

func (p *Panel) HandlePointerLeave() {
	p.lock.Lock()
	defer p.lock.Unlock()

	if top, exists := p.windows.Top(); exists {
		top.HandlePointerLeave()
	}
}

func (p *Panel) AckFrame() {

}
