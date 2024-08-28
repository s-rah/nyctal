// Read input events from evdev (Linux input device subsystem)
package evdev

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

// https://github.com/gvalkov/golang-evdev implements what we do, but requires Cgo

/*
	Hunting the numeric code for EVIOCGRAB was an absolute pain (none of these specify it legibly):

- https://github.com/gvalkov/golang-evdev/search?q=EVIOCGRAB&unscoped_q=EVIOCGRAB
- https://github.com/pfirpfel/node-exclusive-keyboard/blob/master/lib/eviocgrab.cpp
- https://docs.rs/ioctl/0.3.1/src/ioctl/.cargo/registry/src/github.com-1ecc6299db9ec823/ioctl-0.3.1/src/platform/linux.rs.html#396
- https://github.com/torvalds/linux/blob/4a3033ef6e6bb4c566bd1d556de69b494d76976c/include/uapi/linux/input.h#L183
*/
const (
	EVIOCGRAB = 1074021776 // found out with help of evdev.EVIOCGRAB (github.com/gvalkov/golang-evdev)
)

// Sidenote: code for enumerating input devices: https://github.com/MarinX/keylogger/blob/master/keylogger.go

func NewChan() chan InputEvent {
	return make(chan InputEvent)
}

type Device struct {
	Input  chan InputEvent
	handle *os.File
}

func Open(dev string) (*Device, func() error, error) {
	return OpenWithChan(dev, NewChan())
}

func OpenWithChan(dev string, ch chan InputEvent) (*Device, func() error, error) {
	handle, err := os.OpenFile(dev, os.O_RDWR, 0700)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return &Device{
		Input:  ch,
		handle: handle,
	}, handle.Close, nil
}

// same as ScanInput() but grabbed means exclusive access (we'll be the only one receiving the events)
// https://stackoverflow.com/a/1698686
// https://stackoverflow.com/a/1550320
func (d *Device) ScanInputGrabbed(ctx context.Context) error {
	// other programs won't receive input while we have this file handle open
	if err := d.Grab(); err != nil {
		return err
	}

	return d.ScanInput(ctx)
}

func (d *Device) Grab() error {
	return grabExclusiveInputDeviceAccess(d.handle, true)
}

func (d *Device) Ungrab() error {
	return grabExclusiveInputDeviceAccess(d.handle, false)
}

// ScanInput() may close the given channel
func (d *Device) ScanInput(ctx context.Context) error {
	readingErrored := make(chan error, 1)

	go func() {
		for {
			e, err := readOneInputEvent(d.handle)
			if err != nil {
				readingErrored <- fmt.Errorf("readOneInputEvent: %w", err)
				close(d.Input)
				break
			}

			if e != nil {
				d.Input <- *e
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			// this triggers close of inputDevice which will result in return of readOneInputEvent()
			// whose error will be returned to readingErrored, but that's ok because it's buffered
			// and will now not be read because we just exited gracefully
			return nil
		case err := <-readingErrored:
			return err
		}
	}
}

func (d *Device) SetLedOn(led Led) error {
	return d.writeEvent(EvLed, uint16(led), 0x01)
}

func (d *Device) SetLedOff(led Led) error {
	return d.writeEvent(EvLed, uint16(led), 0x00)
}

func (d *Device) writeEvent(evType EventType, code uint16, value int32) error {
	event := &InputEvent{
		// TODO: we're sending time as zero?
		Type:  evType,
		Code:  code,
		Value: value,
	}
	_, err := d.handle.Write(event.AsBytes())
	return err
}

func readOneInputEvent(inputDevice *os.File) (*InputEvent, error) {
	buffer := make([]byte, eventsize)
	n, err := inputDevice.Read(buffer)
	if err != nil {
		return nil, err
	}
	// no input, dont send error
	if n <= 0 {
		return nil, nil
	}
	return InputEventFromBytes(buffer)
}

func grabExclusiveInputDeviceAccess(inputDevice *os.File, grab bool) error {
	// 1 for grab, 0 for un-grab
	grabNum := func() int {
		if grab {
			return 1
		} else {
			return 0
		}
	}()

	if err := unix.IoctlSetInt(int(inputDevice.Fd()), EVIOCGRAB, grabNum); err != nil {
		return fmt.Errorf("grabExclusiveInputDeviceAccess: IOCTL(EVIOCGRAB): %w", err)
	}

	return nil
}

func InputEventFromBytes(buffer []byte) (*InputEvent, error) {
	event := &InputEvent{}
	err := binary.Read(bytes.NewBuffer(buffer), binary.LittleEndian, event)
	return event, err
}
