package evdev

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

// https://python-evdev.readthedocs.io/en/latest/apidoc.html
const (
	KeyRelease = 0
	KeyPress   = 1
	KeyHold    = 2
)

// eventsize is size of structure of InputEvent
var eventsize = int(unsafe.Sizeof(InputEvent{}))

// Has to match the C struct/wire protocol
type InputEvent struct {
	Time  syscall.Timeval
	Type  EventType // EvKey | EvSyn | EvLed | ...
	Code  uint16    // Code is the actual key code, e.g. KeyLEFTSHIFT
	Value int32     // additional specifier like Press/Hold/Release when Type=EvKey
}

func (i *InputEvent) TimevalToTime() time.Time {
	sec, nsec := i.Time.Unix()
	return time.Unix(sec, nsec)
}

func (i *InputEvent) String() string {
	switch i.Type {
	case EvSyn:
		return "SYN"
	case EvKey:
		keyStr := KeyOrButton(i.Code).String()
		switch i.Value {
		case KeyPress:
			return fmt.Sprintf("KeyPress: %s", keyStr)
		case KeyHold:
			return fmt.Sprintf("KeyHold: %s", keyStr)
		case KeyRelease:
			return fmt.Sprintf("KeyRelease: %s", keyStr)
		default:
			return fmt.Sprintf("KeyUNKNOWN(%d): %s", i.Value, keyStr)
		}
	case EvRel:
		return fmt.Sprintf("REL %s %d", Rel(i.Code).String(), i.Value)
	case EvAbs:
		return fmt.Sprintf("ABS %s %d", Abs(i.Code).String(), i.Value)
	case EvMsc:
		return "MSC"
	case EvSw:
		return "SW"
	case EvLed:
		return "LED"
	case EvSnd:
		return "SND"
	case EvRep:
		return "REP"
	case EvFf:
		return "FF"
	case EvPwr:
		return "PWR"
	case EvFfStatus:
		return "FFSTATUS"
	default:
		return fmt.Sprintf("unknown: %d", i.Type)
	}
}

// marshals to evdev "wire format"
func (i *InputEvent) AsBytes() []byte {
	buf := bytes.NewBuffer(make([]byte, 0, 24)) // 24 = sizeof(InputEvent)
	if err := binary.Write(buf, binary.LittleEndian, i); err != nil {
		panic(err)
	}
	return buf.Bytes()
}
