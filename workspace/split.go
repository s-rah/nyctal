package workspace

import (
	"image"
	"image/color"
	"nyctal/model"
	"nyctal/utils"
	"sync"
)

type SplitPanel struct {
	first           model.Workspace
	second          model.Workspace
	splitAt         float64
	activeSplit     bool
	splitHorizontal bool
	bounds          image.Rectangle
	focusSplit      model.Workspace
	do              *DragOverlay
	lock            sync.Mutex
}

func NewSplitPanel(first model.Workspace, do *DragOverlay) *SplitPanel {
	return &SplitPanel{do: do, activeSplit: false, first: first, focusSplit: first}
}

func (p *SplitPanel) AddTopLevel(window model.TopLevelWindow) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.focusSplit.AddTopLevel(window)
}

func (p *SplitPanel) GetTopLevel(idx model.GlobalIdx) model.TopLevelWindow {
	p.lock.Lock()
	defer p.lock.Unlock()
	first := p.first.GetTopLevel(idx)
	if p.activeSplit && first == nil {
		return p.second.GetTopLevel(idx)
	}
	return first
}

func (p *SplitPanel) RemoveTopLevel(idx model.GlobalIdx) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.first.RemoveTopLevel(idx)
	if p.activeSplit {
		p.second.RemoveTopLevel(idx)
	}
}

func (p *SplitPanel) RemoveAllWithParent(pidx model.GlobalIdx) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.first.RemoveAllWithParent(pidx)
	if p.activeSplit {
		p.second.RemoveAllWithParent(pidx)
	}
}

func (p *SplitPanel) Buffer(img *model.BGRA, width int, height int) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.bounds = image.Rect(0, 0, width, height)
	if p.activeSplit {
		firstBounds, secondBounds := p.getBounds()
		// we need to adjust the sub image bounds to account for our split offset...
		// SubImage should really just do this by itself...
		p.first.Buffer(img.SubImage(firstBounds.Add(img.Bounds().Min)), firstBounds.Dx(), firstBounds.Dy())
		p.second.Buffer(img.SubImage(secondBounds.Add(img.Bounds().Min)), secondBounds.Dx(), secondBounds.Dy())

		bounds := firstBounds.Add(img.Bounds().Min)
		if p.splitHorizontal {
			img.DrawRect(bounds.Min.X, bounds.Min.Y+firstBounds.Dy(), bounds.Min.X+firstBounds.Dx(), bounds.Min.Y+firstBounds.Dy(), color.RGBA{R: 255, G: 255, B: 255})
		} else {
			img.DrawRect(bounds.Min.X+firstBounds.Dx(), bounds.Min.Y, bounds.Min.X+firstBounds.Dx(), bounds.Min.Y+firstBounds.Dy(), color.RGBA{R: 255, G: 255, B: 255})
		}
	} else {
		p.first.Buffer(img, width, height)
	}
}

func (p *SplitPanel) getBounds() (image.Rectangle, image.Rectangle) {
	if p.splitHorizontal {
		ypoint := int(p.splitAt * float64(p.bounds.Dy()))
		leftBounds := image.Rect(0, 0, p.bounds.Dx(), ypoint)
		rightBounds := image.Rect(0, ypoint, p.bounds.Dx(), p.bounds.Dy())
		return leftBounds, rightBounds

	}
	xpoint := int(p.splitAt * float64(p.bounds.Dx()))
	leftBounds := image.Rect(0, 0, xpoint, p.bounds.Dy())
	rightBounds := image.Rect(xpoint, 0, p.bounds.Dx(), p.bounds.Dy())
	return leftBounds, rightBounds
}

func (p *SplitPanel) ProcessKeyboardEvent(pointer model.Pointer, kb model.Keyboard, ev model.KeyboardEvent) {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.activeSplit {
		downKeys := kb.DownKeys()
		if downKeys[model.KB_CTRL] && downKeys[model.KB_ALT] && downKeys[36] {
			p.splitAt -= 0.1
		} else if downKeys[model.KB_CTRL] && downKeys[model.KB_ALT] && downKeys[38] {
			p.splitAt += 0.1
		} else {

			fistBounds, secondBounds := p.getBounds()
			if ptr := pointer.ToLocalPointer(fistBounds); ptr != nil {
				p.first.ProcessKeyboardEvent(*ptr, kb, ev)
			} else if ptr := pointer.ToLocalPointer(secondBounds); ptr != nil {
				p.second.ProcessKeyboardEvent(*ptr, kb, ev)
			}
		}

	} else {
		downKeys := kb.DownKeys()

		if downKeys[model.KB_CTRL] && downKeys[model.KB_ALT] && downKeys[47] { // ctrl-alt-v
			utils.Debug(0, "splitpanel", "splitting vertically")
			p.activeSplit = true
			p.splitAt = 0.5
			prev := p.first
			p.first = NewSplitPanel(prev, p.do)
			p.second = NewSplitPanel(NewWindowPanel(p.do), p.do)
			p.splitHorizontal = false
		} else if downKeys[model.KB_CTRL] && downKeys[model.KB_ALT] && downKeys[35] { // ctrl-alt-h
			utils.Debug(0, "splitpanel", "splitting horizontally")
			p.activeSplit = true
			p.splitAt = 0.5
			prev := p.first
			p.first = NewSplitPanel(prev, p.do)
			p.second = NewSplitPanel(NewWindowPanel(p.do), p.do)
			p.splitHorizontal = true
			p.focusSplit = p.first
		} else {
			p.first.ProcessKeyboardEvent(pointer, kb, ev)
		}
	}
}

func (p *SplitPanel) ProcessPointerEvent(pointer model.Pointer, kb model.Keyboard, ev model.PointerEvent) bool {
	p.lock.Lock()
	defer p.lock.Unlock()
	downKeys := kb.DownKeys()
	if downKeys[model.KB_CTRL] && downKeys[model.KB_ALT] {
		if ev.Button != nil {
			if ev.Button.Button == 0x110 && ev.Button.State == 0x00 && p.do.dragging != nil {
				if p.activeSplit {
					fistBounds, secondBounds := p.getBounds()
					if ptr := pointer.ToLocalPointer(fistBounds); ptr != nil {
						window := p.do.dragging
						p.first.AddTopLevel(window)
						p.do.Stop()
						p.focusSplit = p.first
					} else if ptr := pointer.ToLocalPointer(secondBounds); ptr != nil {
						window := p.do.dragging
						p.second.AddTopLevel(window)
						p.do.Stop()
						p.focusSplit = p.second
					}
				} else {
					window := p.do.dragging
					p.first.AddTopLevel(window)
					p.do.Stop()
					p.focusSplit = p.first
				}
				return true
			}
		}
	}

	if p.activeSplit {
		fistBounds, secondBounds := p.getBounds()
		if ptr := pointer.ToLocalPointer(fistBounds); ptr != nil {
			p.first.ProcessPointerEvent(*ptr, kb, ev)
			p.second.HandlePointerLeave()
			p.focusSplit = p.first
		} else if ptr := pointer.ToLocalPointer(secondBounds); ptr != nil {
			p.second.ProcessPointerEvent(*ptr, kb, ev)
			p.first.HandlePointerLeave()
			p.focusSplit = p.second
		}

	} else {
		p.focusSplit = p.first
		return p.first.ProcessPointerEvent(pointer, kb, ev)
	}
	return false
}

func (p *SplitPanel) HandlePointerLeave() {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.activeSplit {
		p.first.HandlePointerLeave()
		p.second.HandlePointerLeave()

	} else {
		p.first.HandlePointerLeave()
	}
}
