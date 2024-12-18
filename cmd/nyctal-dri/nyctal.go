package main

import (
	"context"
	"fmt"
	"image"
	"os"
	"os/exec"

	//"os/exec"
	"runtime/debug"
	"time"

	"nyctal-dri/evdev"
	"nyctal/model"
	"nyctal/utils"
	"nyctal/wayland"
	"nyctal/workspace"
)

var POINTER model.Pointer
var KEYBOARD = model.NewKeyboardModel()

func SetupMouse(workspace model.Workspace) func() error {

	mouseDev := evdev.FindMouseDevice()
	if mouseDev != "" {
		dev, f, err := evdev.Open(mouseDev)
		if err != nil {
			fmt.Printf("[error] %s", err)
		} else {
			go func() {
				err := dev.ScanInput(context.Background())
				utils.Debug(0, "input handler", fmt.Sprintf("failed to scan input: %v", err))
			}()
			go func() {
				utils.Debug(0, "mouse handler", "starting input handler")
				localX := float32(0.0)
				localY := float32(0.0)
				for {
					//	utils.Debug("input handler", "waiting...")
					ev := <-dev.Input
					utils.Debug(0, "mouse handler", fmt.Sprintf("type %d code %d", ev.Type, ev.Code))
					switch ev.Type {
					case evdev.EvKey:
						pev := model.PointerEvent{
							Button: &model.PointerButtonEvent{Time: uint32(time.Now().UnixMilli()), Button: uint32(ev.Code), State: uint32(ev.Value)}}
						POINTER.ProcessPointerEvent(pev)
						workspace.ProcessPointerEvent(POINTER, *KEYBOARD, pev)
					case evdev.EvRel:
						if ev.Code == 0x00 {
							// abs x
							localX += float32(ev.Value)
							if localX < 0 {
								localX = 0
							}
							pev := model.PointerEvent{
								Move: &model.PointerMoveEvent{Time: uint32(time.Now().UnixMilli()), MX: localX, MY: localY},
							}
							POINTER.ProcessPointerEvent(pev)
							workspace.ProcessPointerEvent(POINTER, *KEYBOARD, pev)
						}
						if ev.Code == 0x01 {
							// abs y
							localY += float32(ev.Value)
							if localY < 0 {
								localY = 0
							}
							pev := model.PointerEvent{
								Move: &model.PointerMoveEvent{Time: uint32(time.Now().UnixMilli()), MX: localX, MY: localY},
							}
							POINTER.ProcessPointerEvent(pev)
							workspace.ProcessPointerEvent(POINTER, *KEYBOARD, pev)
						}
						if ev.Code == 0x08 {
							pev := model.PointerEvent{
								Axis: &model.PointerAxisEvent{Time: uint32(time.Now().UnixMilli()), Value: float32(ev.Value)},
							}
							POINTER.ProcessPointerEvent(pev)
							workspace.ProcessPointerEvent(POINTER, *KEYBOARD, pev)
						}
					}

				}
			}()
			return f
		}

	} else {
		utils.Debug(0, "input handler", "coult not find mouse device")
	}
	return func() error { return nil }
}

// see include/uapi/linux/input-event-codes.h
func SetupInput(workspace model.Workspace) func() error {

	kdev := evdev.FindAllKeyboardDevices()
	if len(kdev) > 0 {
		dev, f, err := evdev.Open(kdev[0])
		if err != nil {
			fmt.Printf("[error] %s", err)
		} else {
			go func() {
				err := dev.ScanInput(context.Background())
				utils.Debug(0, "input handler", fmt.Sprintf("failed to scan input: %v", err))
			}()
			go func() {
				utils.Debug(0, "input handler", "starting input handler")

				for {
					//	utils.Debug("input handler", "waiting...")
					ev := <-dev.Input
					utils.Debug(0, "input handler", fmt.Sprintf("type %d code %d value: %d", ev.Type, ev.Code, ev.Value))
					kev := model.KeyboardEvent{Time: uint32(time.Now().UnixMilli()), Key: uint32(ev.Code), State: uint32(ev.Value)}
					KEYBOARD.ProcessKeyboardEvent(kev)
					workspace.ProcessKeyboardEvent(POINTER, *KEYBOARD, kev)
				}
			}()
			return f
		}
	}
	return func() error { return nil }
}

func main() {

	debug.SetPanicOnFault(false)

	var output model.Output
	wspace := workspace.NewDragOverlay()
	width, height := uint32(1024), uint32(1024)
	rm := DrmInit()
	if rm == nil {
		fmt.Printf("[error] could not initialize drm rendering\n")
		output = NewImageOutput("nyctal-", time.Second*5)
	} else {
		output = rm
		width, height = rm.Stats()
		wspace = workspace.NewDragOverlay()
		utils.Debug(0, "nyctal", fmt.Sprintf("dri established %v %v", width, height))
	}

	// reset UDS
	os.RemoveAll("/tmp/nyctal/")
	os.Mkdir("/tmp/nyctal/", 0700)

	ws, err := wayland.NewServer("/tmp/nyctal/nyctal-0", wspace)
	if err != nil {
		fmt.Printf("[error] %s\n", err)
		os.Exit(1)
	}
	go ws.Listen()
	closeInput := SetupInput(wspace)
	defer closeInput()
	closeMInput := SetupMouse(wspace)
	defer closeMInput()

	fmt.Printf("Starting Nyctal...\n")
	lastFrame := time.Now()

	cmd := exec.Command("/bin/elope")
	go func() {
		time.Sleep(time.Second * 2)
		stdoutStderr, err := cmd.CombinedOutput()
		if err != nil {
		}
		fmt.Printf("%s\n", stdoutStderr)
		os.Exit(1)
	}()

	for {
		// time check is here to prevent spamming the workspace render buffer (which may attempt to e.g. reconfigure windwows)
		// there is no point in attempting to generate frames any faster than 200fps
		// todo: in the future we should replace this with a NeedsRender() check
		if time.Since(lastFrame) >= time.Millisecond*5 {
			buffer := model.EmptyBGRA(image.Rect(0, 0, int(width), int(height)))
			wspace.Buffer(buffer, int(width), int(height))
			output.RenderBuffer(buffer)
			lastFrame = time.Now()
		}
	}
}
