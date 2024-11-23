package model

import (
	"image"
	"nyctal/utils"
	"time"
)

type Pointer struct {
	MX   int
	MY   int
	OX   int
	OY   int
	Time time.Time
}

func (p *Pointer) ProcessPointerEvent(ev PointerEvent) bool {
	if ev.Move != nil && time.Since(p.Time).Milliseconds() > 1 {
		p.MX = int(ev.Move.MX)
		p.MY = int(ev.Move.MY)
		p.OX = int(ev.Move.MX)
		p.OY = int(ev.Move.MY)
		p.Time = time.Now()
		return true
	}
	return false
}

func (p *Pointer) ToLocalPointer(rect image.Rectangle) *Pointer {
	lp := &Pointer{MX: p.MX, MY: p.MY, OX: p.OX, OY: p.OY}
	if intersects, lx, ly := utils.IntersectsRect(int(p.MX), int(p.MY), rect); intersects {
		lp.MX = lx
		lp.MY = ly
		return lp
	}
	return nil
}
