package main

/*
#cgo CFLAGS: -I ./minifb/include/
#cgo linux LDFLAGS: -L./minifb/build/ -lminifb -lX11 -lGL -lXfixes
#include <MiniFB.h>
#include <stdlib.h>

// callback exported from Go
extern void Keyboard(struct mfb_window*, mfb_key, mfb_key_mod, bool);
extern void MouseButton(struct mfb_window*, mfb_key, mfb_key_mod, bool);
extern void MouseMove(struct mfb_window*, int, int);
extern void MouseScroll(struct mfb_window*, mfb_key_mod, float, float);
extern void ResizeWindow(struct mfb_window*, int, int);
*/
import "C"
import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/color"
	"sync"

	"os"
	"runtime"

	"runtime/pprof"
	"time"
	"unsafe"

	"nyctal/model"
	"nyctal/wayland"
	"nyctal/workspace"
)

var window *C.struct_mfb_window

func create_window(title string, w, h int) error {
	ctitle := C.CString(title)
	cw := C.uint(w)
	ch := C.uint(h)

	// ctitle, window and buffer are all dynamically allocated by C so we could free them, but they live for the entirety of the program so I didn't bother
	// title, width, height, flags
	window = C.mfb_open_ex(ctitle, cw, ch, C.WF_RESIZABLE)
	if window == nil {
		return errors.New("could not create window")
	}

	// allocate pixel buffer
	ResizeWindow(window, C.int(w), C.int(h))
	return nil
}

//export MouseButton
func MouseButton(window *C.struct_mfb_window, button C.mfb_key, mod C.mfb_key_mod, isPressed C.bool) {
	//x := C.mfb_get_mouse_x(window)
	//	y := C.mfb_get_mouse_y(window)
	state := uint32(0)
	if isPressed {
		state = 1
	}
	wspace.ProcessPointerEvent(POINTER, *KEYBOARD, model.PointerEvent{Button: &model.PointerButtonEvent{Time: uint32(time.Now().UnixMilli()), Button: uint32(button - 1 + 0x110), State: state}})
}

//export MouseScroll
func MouseScroll(window *C.struct_mfb_window, mod C.mfb_key_mod, dx C.float, dy C.float) {
	wspace.ProcessPointerEvent(POINTER, *KEYBOARD, model.PointerEvent{Axis: &model.PointerAxisEvent{Time: uint32(time.Now().UnixMilli()), Axis: 0, Value: float32(-dy * 4.0)}})
}

//export MouseMove
func MouseMove(window *C.struct_mfb_window, mx C.int, my C.int) {
	x := C.mfb_get_mouse_x(window)
	y := C.mfb_get_mouse_y(window)
	if x > 0 && y > 0 {
		ev := model.PointerEvent{Move: &model.PointerMoveEvent{Time: uint32(time.Now().UnixMilli()), MX: float32(int(x)), MY: float32(int(y))}}
		POINTER.ProcessPointerEvent(ev)
		wspace.ProcessPointerEvent(POINTER, *KEYBOARD, ev)
	}
}

var running = true

//export Keyboard
func Keyboard(window *C.struct_mfb_window, key C.mfb_key, mod C.mfb_key_mod, isPressed C.bool) {

	switch key {
	case C.KB_KEY_ESCAPE:
		running = false
		fmt.Printf("closing...%v\n", time.Now())
		return
	}

	pressed := 0
	if isPressed {
		pressed = 1
	}

	// the keys we get from minifb have already been mapped...

	ev := model.KeyboardEvent{
		Time:  uint32(time.Now().UnixMilli()),
		Key:   uint32(key - 8),
		State: uint32(pressed),
	}
	KEYBOARD.ProcessKeyboardEvent(ev)
	wspace.ProcessKeyboardEvent(POINTER, *KEYBOARD, ev)
}

var buffer *C.uint
var buffer_len int

var wspace model.Workspace
var lock sync.Mutex

var WIDTH int
var HEIGHT int

var POINTER model.Pointer
var KEYBOARD = model.NewKeyboardModel()

//export ResizeWindow
func ResizeWindow(window *C.struct_mfb_window, width C.int, height C.int) {
	lock.Lock()
	defer lock.Unlock()
	w, h := int(width), int(height)
	new_buffer_len := w * h * 4

	buffer = (*C.uint)(C.realloc(unsafe.Pointer(buffer), C.ulong(new_buffer_len)))
	buffer_len = new_buffer_len
	WIDTH = w
	HEIGHT = h
}

// helper function to convert a image/color to a C.uint used by minifb's buffer
func value(col color.Color) C.uint {
	if c, ok := col.(color.RGBA); ok {
		return C.uint((int(c.R) << 16) | (int(c.G) << 8) | int(c.B))
	} else {
		return 0x00
	}
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			fmt.Printf("[error] %s", err)
			os.Exit(1)
		}
		pprof.StartCPUProfile(f)
		defer func() {
			pprof.StopCPUProfile()
			f, _ := os.Create("mem.prof")
			pprof.WriteHeapProfile(f)

		}()
	}

	runtime.LockOSThread()

	wspace = workspace.NewDragOverlay()

	e := create_window("nyctal-x11", 128, 128)

	if e != nil {
		panic(e)
	}

	// setup keyboard handler
	// notice that I pass the C.Keyboard callback here casted to the C.mfb_keyboard_func type
	// C.Keyboard is a C function but it was implemented in Go, this is the easiest way to pass go callbacks to C
	C.mfb_set_keyboard_callback(window, (C.mfb_keyboard_func)(C.Keyboard))
	C.mfb_set_mouse_button_callback(window, (C.mfb_mouse_button_func)(C.MouseButton))
	C.mfb_set_mouse_move_callback(window, (C.mfb_mouse_move_func)(C.MouseMove))
	C.mfb_set_mouse_scroll_callback(window, (C.mfb_mouse_scroll_func)(C.MouseScroll))
	C.mfb_set_resize_callback(window, (C.mfb_resize_func)(C.ResizeWindow))

	// reset UDS
	os.RemoveAll("/tmp/nyctal/")
	os.Mkdir("/tmp/nyctal/", 0700)

	ws, err := wayland.NewServer("/tmp/nyctal/nyctal-0", wspace)
	if err != nil {
		fmt.Printf("[error] %s\n", err)
		os.Exit(1)
	}
	go ws.Listen()
	//wspace.ProcessFocus()

	fmt.Printf("Starting nyctal-x11...\n")
	//lastFrame := time.Now()
	// this is the minifb loop

	for C.mfb_wait_sync(window) {

		// // // this will be set to false if the esc key is pressed
		// if ws, ok := wspace.(*workspace.Workspace); ok {
		// 	if ws.Quit {
		// 		fmt.Printf("closing...%v\n", time.Now())
		// 		C.mfb_close(window)
		// 		break
		// 	}
		// }

		img := model.EmptyBGRA(image.Rect(0, 0, WIDTH, HEIGHT))
		wspace.Buffer(img, WIDTH, HEIGHT)
		bounds := img.Bounds()
		w, h := bounds.Max.X, bounds.Max.Y
		if w*h*4 == buffer_len {
			pixels := unsafe.Slice((*C.uint)(buffer), w*h)
			for j := 0; j < h; j++ {
				for i := 0; i < w; i++ {
					pixels[(w*j)+i] = value(img.AtRaw(i, j))
				}
			}

			// minifb stuff
			state := C.mfb_update_ex(window, unsafe.Pointer(buffer), C.uint(w), C.uint(h))
			if state < 0 {
				break
			}
		}

		// time check is here to prevent spamming the workspace render buffer (which may attempt to e.g. reconfigure windwows)
		// there is no point in attempting to generate frames any faster than 200fps
		// todo: in the future we should replace this with a NeedsRender() check
		//if time.Since(lastFrame) >= time.Millisecond*5 {
		//wspace.AckFrame()
		//lastFrame = time.Now()
		//}
	}
}
