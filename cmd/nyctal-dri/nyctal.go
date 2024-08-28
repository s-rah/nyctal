package main

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"nyctal-dri/evdev"
	"nyctal/model"
	"nyctal/utils"
	"nyctal/wayland"
	"nyctal/workspace"
)

// see include/uapi/linux/input-event-codes.h
func SetupInput(workspace model.Client) func() error {

	kdev := evdev.FindAllKeyboardDevices()
	if len(kdev) > 0 {
		dev, f, err := evdev.Open(kdev[0])
		if err != nil {
			fmt.Printf("[error] %s", err)
		} else {
			go func() {
				err := dev.ScanInput(context.Background())
				utils.Debug("input handler", fmt.Sprintf("failed to scan input: %v", err))
			}()
			go func() {
				utils.Debug("input handler", "starting input handler")

				for {
					//	utils.Debug("input handler", "waiting...")
					ev := <-dev.Input
					utils.Debug("input handler", fmt.Sprintf("type %d code %d value: %d", ev.Type, ev.Code, ev.Value))

				}
			}()
			return f
		}
	}

	mouseDev := evdev.FindMouseDevice()
	if mouseDev != "" {
		dev, f, err := evdev.Open(mouseDev)
		if err != nil {
			fmt.Printf("[error] %s", err)
		} else {
			go func() {
				err := dev.ScanInput(context.Background())
				utils.Debug("input handler", fmt.Sprintf("failed to scan input: %v", err))
			}()
			go func() {
				utils.Debug("input handler", "starting input handler")
				localX := float32(0.0)
				localY := float32(0.0)
				for {
					//	utils.Debug("input handler", "waiting...")
					ev := <-dev.Input
					//utils.Debug("input handler", fmt.Sprintf("type %d code %d", ev.Type, ev.Code))
					switch ev.Type {
					case evdev.EvKey:
						workspace.ProcessPointerEvent(model.PointerEvent{
							Button: &model.PointerButtonEvent{Time: uint32(time.Now().UnixMilli()), Button: uint32(ev.Code), State: uint32(ev.Value)},
						})
					case evdev.EvRel:
						if ev.Code == 0x00 {
							// abs x
							localX += float32(ev.Value)
							if localX < 0 {
								localX = 0
							}
							workspace.ProcessPointerEvent(model.PointerEvent{
								Move: &model.PointerMoveEvent{Time: uint32(time.Now().UnixMilli()), MX: localX, MY: localY},
							})
						}
						if ev.Code == 0x01 {
							// abs y
							localY += float32(ev.Value)
							if localY < 0 {
								localY = 0
							}
							workspace.ProcessPointerEvent(model.PointerEvent{
								Move: &model.PointerMoveEvent{Time: uint32(time.Now().UnixMilli()), MX: localX, MY: localY},
							})
						}
						if ev.Code == 0x08 {
							workspace.ProcessPointerEvent(model.PointerEvent{
								Axis: &model.PointerAxisEvent{Time: uint32(time.Now().UnixMilli()), Value: float32(ev.Value)},
							})
						}
					}

				}
			}()
			return f
		}

	} else {
		utils.Debug("input handler", "coult not find mouse device")
	}
	return func() error { return nil }
}

func main() {

	debug.SetPanicOnFault(false)

	var output model.Output
	wspace := workspace.NewWorkspace(1280, 1024)

	rm := DrmInit()
	if rm == nil {
		fmt.Printf("[error] could not initialize drm rendering\n")
		output = NewImageOutput("ori-", time.Second*5)
	} else {
		output = rm
		width, height := rm.Stats()
		wspace = workspace.NewWorkspace(int(width), int(height))
		utils.Debug("ori", fmt.Sprintf("dri established %v %v", width, height))
	}

	// reset UDS
	os.RemoveAll("/tmp/ori/")
	os.Mkdir("/tmp/ori/", 0700)

	ws, err := wayland.NewServer("/tmp/ori/ori-0", wspace)
	if err != nil {
		fmt.Printf("[error] %s\n", err)
		os.Exit(1)
	}
	go ws.Listen()
	closeInput := SetupInput(wspace)
	defer closeInput()
	wspace.ProcessFocus()

	fmt.Printf("Starting ORI...\n")
	lastFrame := time.Now()
	start := time.Now()
	start = start.Add(-time.Second * 10)
	for {
		// time check is here to prevent spamming the workspace render buffer (which may attempt to e.g. reconfigure windwows)
		// there is no point in attempting to generate frames any faster than 200fps
		// todo: in the future we should replace this with a NeedsRender() check
		if time.Since(lastFrame) >= time.Millisecond*5 {
			buffer := wspace.Buffer()
			output.RenderBuffer(buffer)
			wspace.AckFrame()
			lastFrame = time.Now()
		}
	}
}
