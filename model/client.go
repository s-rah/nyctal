package model

import ()

type PointerEvent struct {
	Move   *PointerMoveEvent
	Button *PointerButtonEvent
	Axis   *PointerAxisEvent
}

type PointerMoveEvent struct {
	Time uint32
	MX   float32
	MY   float32
}

type PointerButtonEvent struct {
	Time   uint32
	Button uint32
	State  uint32
}

type PointerAxisEvent struct {
	Time  uint32
	Axis  uint32
	Value float32
}

type KeyboardEvent struct {
	Time      uint32
	Key       uint32
	State     uint32
	Modifiers uint32
}

type Buffer struct {
	Data   []byte
	Offset int
	Width  int
	Height int
	Stride int
	Format Format
}

type Client interface {
	Parent() Client
	AddChild(id int, client Client) int
	RemoveChild(id int)
	RemoveChildren(parent int)
	GetChild(id int) Client
	Resize(width int, height int)
	Buffer() *BGRA

	ProcessKeyboardEvent(ev KeyboardEvent)
	ProcessPointerEvent(ev PointerEvent) bool

	HandlePointerLeave()

	ProcessFocus()
	ProcessUnFocus()

	AckFrame()
}
