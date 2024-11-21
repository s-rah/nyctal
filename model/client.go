package model

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

type TopLevelWindow interface {
	Index() GlobalIdx
	Parent() GlobalIdx
	Buffer(img *BGRA, width int, height int)
	ProcessKeyboardEvent(ev KeyboardEvent)
	ProcessPointerEvent(ev PointerEvent) bool
	HandlePointerLeave()
	AckFrame()
}

type Workspace interface {
	AddTopLevel(TopLevelWindow)
	RemoveTopLevel(GlobalIdx)
	RemoveAllWithParent(GlobalIdx)
	GetTopLevel(GlobalIdx) TopLevelWindow

	Buffer(img *BGRA, width int, height int)

	ProcessKeyboardEvent(pointer Pointer, kb Keyboard, ev KeyboardEvent)
	ProcessPointerEvent(pointer Pointer, kb Keyboard, ev PointerEvent) bool
	HandlePointerLeave()
}
