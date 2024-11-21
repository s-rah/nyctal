package model

import (
	"fmt"
	"image"
	"image/color"

	"nyctal/utils"
)

type Format int

const (
	FormatARGB = Format(0)
	FormatXRGB = Format(1)
)

// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

type BGRA struct {
	Pix    []byte
	Rect   image.Rectangle
	Stride int
}

func EmptyBGRA(size image.Rectangle) *BGRA {
	return &BGRA{
		Pix:    make([]byte, size.Dy()*size.Dx()*4),
		Rect:   size,
		Stride: size.Dx() * 4,
	}
}

func NewBGRA(pixels []byte, rect image.Rectangle, stride int) *BGRA {
	pixelLen := rect.Dy() * stride
	cpix := make([]byte, pixelLen)
	copy(cpix[0:pixelLen], pixels[0:pixelLen])
	return &BGRA{
		Pix:    cpix,
		Rect:   rect,
		Stride: stride,
	}
}

// HLine draws a horizontal line
func (i *BGRA) HLine(x1, y, x2 int, col color.RGBA) {
	for ; x1 <= x2; x1++ {
		i.Set(x1, y, col)
	}
}

// VLine draws a veritcal line
func (i *BGRA) VLine(x, y1, y2 int, col color.RGBA) {
	for ; y1 <= y2; y1++ {
		i.Set(x, y1, col)
	}
}

// Rect draws a rectangle utilizing HLine() and VLine()
func (i *BGRA) DrawRect(x1, y1, x2, y2 int, col color.RGBA) {
	i.HLine(x1, y1, x2, col)
	i.HLine(x1, y2, x2, col)
	i.VLine(x1, y1, y2, col)
	i.VLine(x2, y1, y2, col)
}

func (i *BGRA) Update(pixels []byte, damage image.Rectangle, stride int) {

	if i.Bounds().Max.X < damage.Bounds().Min.X {
		utils.Debug("cached image", fmt.Sprintf("OOB damaged buffer: %v %v %v", i.Bounds(), damage, stride))
		return
	}

	minY := min(i.Rect.Bounds().Max.Y, damage.Min.Y)
	maxY := min(i.Rect.Bounds().Max.Y, damage.Max.Y)

	for sy := minY; sy < maxY; sy++ {
		damageY := sy * i.Stride

		damageStartX := min(i.Rect.Bounds().Max.X, damage.Min.X)
		damageEndX := min(i.Rect.Bounds().Max.X, damage.Max.X)

		damageStart := damageY + (damageStartX * 4)
		damageEnd := damageY + (damageEndX * 4)

		copy(i.Pix[damageStart:damageEnd], pixels[damageStart:damageEnd])
	}
	//pixelLen := i.Rect.Dy() * stride
	//copy(i.Pix[0:pixelLen], pixels[0:pixelLen])

}

func (i *BGRA) SubImage(bounds image.Rectangle) *BGRA {
	bounds = bounds.Intersect(i.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if bounds.Empty() {
		return &BGRA{}
	}
	p := i.PixOffset(bounds.Min.X, bounds.Min.Y)
	return &BGRA{
		Pix:    i.Pix[p:],
		Stride: i.Stride,
		Rect:   bounds,
	}
}

func (i *BGRA) Damage(start int, end int, new []byte) {
	copy(i.Pix[start:end], new[start:end])
}

func (i *BGRA) Bounds() image.Rectangle { return i.Rect }
func (i *BGRA) ColorModel() color.Model { return color.RGBAModel }

func (i *BGRA) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(i.Rect)) {
		return color.RGBA{}
	}

	n := i.PixOffset(x, y)
	pix := i.Pix[n:]
	return color.RGBA{pix[2], pix[1], pix[0], pix[3]}
}

func (i *BGRA) AtRaw(x, y int) color.RGBA {
	if !(image.Point{x, y}.In(i.Rect)) {
		return color.RGBA{}
	}

	n := i.PixOffset(x, y)
	pix := i.Pix[n:]
	return color.RGBA{pix[2], pix[1], pix[0], pix[3]}
}

func (i *BGRA) Set(x, y int, c color.Color) {
	i.SetRGBA(x, y, color.RGBAModel.Convert(c).(color.RGBA))
}

func (i *BGRA) SetRGBA(x, y int, c color.RGBA) {
	if !(image.Point{x, y}.In(i.Rect)) {
		return
	}

	n := i.PixOffset(x, y)
	pix := i.Pix[n:]

	pix[0] = c.B
	pix[1] = c.G
	pix[2] = c.R
	pix[3] = c.A
}

func (i *BGRA) PixOffset(x, y int) int {
	return (y-i.Rect.Min.Y)*i.Stride + (x-i.Rect.Min.X)*4
}

const m = 1<<16 - 1

func clip(dst *BGRA, r *image.Rectangle, src *BGRA, sp *image.Point) {
	orig := r.Min
	*r = r.Intersect(dst.Bounds())
	*r = r.Intersect(src.Bounds().Add(orig.Sub(*sp)))
	dx := r.Min.X - orig.X
	dy := r.Min.Y - orig.Y
	if dx == 0 && dy == 0 {
		return
	}
	sp.X += dx
	sp.Y += dy

}

func DrawCopyOver(dst *BGRA, r image.Rectangle, src *BGRA, sp image.Point) {
	clip(dst, &r, src, &sp)
	if r.Empty() {
		return
	}
	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx*4, +4
	} else {
		// If the source start point is higher than the destination start point, or equal height but to the left,
		// then we compose the rows in right-to-left, bottom-up order instead of left-to-right, top-down.
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)*4, -4, -4
	}
	for ; dy > 0; dy-- {
		dpix := dst.Pix[d0:]
		spix := src.Pix[s0:]
		for i := i0; i != i1; i += idelta {
			s := spix[i : i+4 : i+4] // Small cap improves performance, see https://golang.org/issue/27857
			sr := uint32(s[0]) * 0x101
			sg := uint32(s[1]) * 0x101
			sb := uint32(s[2]) * 0x101
			sa := uint32(s[3]) * 0x101

			// The 0x101 is here for the same reason as in drawRGBA.
			a := (m - sa) * 0x101

			d := dpix[i : i+4 : i+4] // Small cap improves performance, see https://golang.org/issue/27857
			d[0] = uint8((uint32(d[0])*a/m + sr) >> 8)
			d[1] = uint8((uint32(d[1])*a/m + sg) >> 8)
			d[2] = uint8((uint32(d[2])*a/m + sb) >> 8)
			d[3] = uint8((uint32(d[3])*a/m + sa) >> 8)
		}
		d0 += ddelta
		s0 += sdelta
	}
}
