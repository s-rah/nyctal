package model

import (
	"image"
	"nyctal/utils"
)

type Pointer struct {
	MX int
	MY int
	OX int
	OY int
}

func (p *Pointer) ProcessPointerEvent(ev PointerEvent) {
	if ev.Move != nil {
		p.MX = int(ev.Move.MX)
		p.MY = int(ev.Move.MY)
		p.OX = int(ev.Move.MX)
		p.OY = int(ev.Move.MY)
	}
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
