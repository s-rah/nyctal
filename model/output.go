package model

type Output interface {
	RenderBuffer(img *BGRA) error
}
