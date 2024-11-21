package workspace

import (
	"fmt"
	"image"
	"image/color"
	"nyctal/model"
	"os"
	"os/exec"
)

type DragOverlay struct {
	SplitPanel
	dragging         model.TopLevelWindow
	startDragPointer model.Pointer
}

func (do *DragOverlay) Start(pointer model.Pointer, window model.TopLevelWindow) {
	do.startDragPointer = pointer
	do.dragging = window
}

func (do *DragOverlay) Update(pointer model.Pointer) {
	do.startDragPointer = pointer
}

func (do *DragOverlay) Stop() {
	do.dragging = nil
}

func (do *DragOverlay) ProcessKeyboardEvent(pointer model.Pointer, kb model.Keyboard, ev model.KeyboardEvent) {
	downKeys := kb.DownKeys()
	if downKeys[model.KB_CTRL] && downKeys[model.KB_ALT] && downKeys[model.KB_ENTER] {
		cmd := exec.Command("./elope")
		cmd.Env = append(cmd.Env, "XDG_RUNTIME_DIR=/tmp/nyctal/", "WAYLAND_DISPLAY=nyctal-0")
		go func() {
			stdoutStderr, err := cmd.CombinedOutput()
			if err != nil {
			}
			fmt.Printf("Error: %s %v\n", stdoutStderr, err)
			os.Exit(1)
		}()
	}
	do.SplitPanel.ProcessKeyboardEvent(pointer, kb, ev)
}

func (do *DragOverlay) ProcessPointerEvent(pointer model.Pointer, kb model.Keyboard, ev model.PointerEvent) bool {
	do.Update(pointer)
	return do.SplitPanel.ProcessPointerEvent(pointer, kb, ev)
}

func (do *DragOverlay) Buffer(img *model.BGRA, width, height int) {
	do.SplitPanel.Buffer(img, width, height)
	if do.dragging != nil {
		do.dragging.Buffer(img.SubImage(image.Rect(do.startDragPointer.OX,
			do.startDragPointer.OY,
			do.startDragPointer.OX+480,
			do.startDragPointer.OY+256)),
			480, 256)
	}
	img.DrawRect(do.startDragPointer.OX, do.startDragPointer.OY, do.startDragPointer.OX+2, do.startDragPointer.OY+2, color.RGBA{R: 255})
}

func NewDragOverlay() model.Workspace {
	do := &DragOverlay{}
	do.SplitPanel = *NewSplitPanel(NewWindowPanel(do), do)
	return do
}
