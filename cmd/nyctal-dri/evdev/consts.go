package evdev

import (
	"fmt"
)

// EventType are groupings of codes under a logical input construct.
// Each type has a set of applicable codes to be used in generating events.
// See the Ev section for details on valid codes for each type
// https://github.com/torvalds/linux/blob/5c8fe583cc/include/uapi/linux/input-event-codes.h#L35
type EventType uint16

const (
	// marker to separate events. Events may be separated in time or in space, such as with the multitouch protocol.
	EvSyn EventType = 0x00
	// state changes of keyboards, buttons, or other key-like devices.
	EvKey EventType = 0x01
	// relative axis value changes, e.g. moving the mouse 5 units to the left.
	EvRel EventType = 0x02
	// absolute axis value changes, e.g. describing the coordinates of a touch on a touchscreen.
	EvAbs EventType = 0x03
	// miscellaneous input data that do not fit into other types.
	EvMsc EventType = 0x04
	// binary state input switches.
	EvSw EventType = 0x05
	// turn LEDs on devices on and off.
	EvLed EventType = 0x11
	// output sound to devices.
	EvSnd EventType = 0x12
	// for autorepeating devices.
	EvRep EventType = 0x14
	// send force feedback commands to an input device.
	EvFf EventType = 0x15
	// special type for power button and switch input.
	EvPwr EventType = 0x16
	// receive force feedback device status.
	EvFfStatus EventType = 0x17
)

// https://github.com/torvalds/linux/blob/5c8fe583cc/include/uapi/linux/input-event-codes.h#L913
type Led uint16

const (
	LedNUML     Led = 0x00
	LedCAPSL    Led = 0x01
	LedSCROLLL  Led = 0x02
	LedCOMPOSE  Led = 0x03
	LedKANA     Led = 0x04
	LedSLEEP    Led = 0x05
	LedSUSPEND  Led = 0x06
	LedMUTE     Led = 0x07
	LedMISC     Led = 0x08
	LedMAIL     Led = 0x09
	LedCHARGING Led = 0x0a
)

// https://github.com/torvalds/linux/blob/5c8fe583cc/include/uapi/linux/input-event-codes.h#L785
type Rel uint16

const (
	RelX           Rel = 0x00
	RelY           Rel = 0x01
	RelZ           Rel = 0x02
	RelRX          Rel = 0x03
	RelRY          Rel = 0x04
	RelRZ          Rel = 0x05
	RelHWHEEL      Rel = 0x06
	RelDIAL        Rel = 0x07
	RelWHEEL       Rel = 0x08
	RelMISC        Rel = 0x09
	RelRESERVED    Rel = 0x0a
	RelWHEELHIRES  Rel = 0x0b
	RelHWHEELHIRES Rel = 0x0c
)

var relToString = map[Rel]string{
	RelX:           "X",
	RelY:           "Y",
	RelZ:           "Z",
	RelRX:          "RX",
	RelRY:          "RY",
	RelRZ:          "RZ",
	RelHWHEEL:      "HWHEEL",
	RelDIAL:        "DIAL",
	RelWHEEL:       "WHEEL",
	RelMISC:        "MISC",
	RelRESERVED:    "RESERVED",
	RelWHEELHIRES:  "WHEEL_HI_RES",
	RelHWHEELHIRES: "HWHEEL_HI_RES",
}

func (r Rel) String() string {
	str, found := relToString[r]
	if !found {
		return fmt.Sprintf("Unknown(%d)", r)
	}

	return str
}

// https://github.com/torvalds/linux/blob/5c8fe583cc/include/uapi/linux/input-event-codes.h#L812
type Abs uint16

const (
	AbsX         Abs = 0x00
	AbsY         Abs = 0x01
	AbsZ         Abs = 0x02
	AbsRX        Abs = 0x03
	AbsRY        Abs = 0x04
	AbsRZ        Abs = 0x05
	AbsTHROTTLE  Abs = 0x06
	AbsRUDDER    Abs = 0x07
	AbsWHEEL     Abs = 0x08
	AbsGAS       Abs = 0x09
	AbsBRAKE     Abs = 0x0a
	AbsHAT0X     Abs = 0x10
	AbsHAT0Y     Abs = 0x11
	AbsHAT1X     Abs = 0x12
	AbsHAT1Y     Abs = 0x13
	AbsHAT2X     Abs = 0x14
	AbsHAT2Y     Abs = 0x15
	AbsHAT3X     Abs = 0x16
	AbsHAT3Y     Abs = 0x17
	AbsPRESSURE  Abs = 0x18
	AbsDISTANCE  Abs = 0x19
	AbsTILTX     Abs = 0x1a
	AbsTILTY     Abs = 0x1b
	AbsTOOLWIDTH Abs = 0x1c

	AbsVOLUME Abs = 0x20

	AbsMISC Abs = 0x28

	AbsRESERVED Abs = 0x2e

	AbsMTSLOT        Abs = 0x2f // MT slot being modified
	AbsMTTOUCHMAJOR  Abs = 0x30 // Major axis of touching ellipse
	AbsMTTOUCHMINOR  Abs = 0x31 // Minor axis (omit if circular)
	AbsMTWIDTHMAJOR  Abs = 0x32 // Major axis of approaching ellipse
	AbsMTWIDTHMINOR  Abs = 0x33 // Minor axis (omit if circular)
	AbsMTORIENTATION Abs = 0x34 // Ellipse orientation
	AbsMTPOSITIONX   Abs = 0x35 // Center X touch position
	AbsMTPOSITIONY   Abs = 0x36 // Center Y touch position
	AbsMTTOOLTYPE    Abs = 0x37 // Type of touching device
	AbsMTBLOBID      Abs = 0x38 // Group a set of packets as a blob
	AbsMTTRACKINGID  Abs = 0x39 // Unique ID of initiated contact
	AbsMTPRESSURE    Abs = 0x3a // Pressure on contact area
	AbsMTDISTANCE    Abs = 0x3b // Contact hover distance
	AbsMTTOOLX       Abs = 0x3c // Center X tool position
	AbsMTTOOLY       Abs = 0x3d // Center Y tool position
)

var absToString = map[Abs]string{
	AbsX:             "X",
	AbsY:             "Y",
	AbsZ:             "Z",
	AbsRX:            "RX",
	AbsRY:            "RY",
	AbsRZ:            "RZ",
	AbsTHROTTLE:      "THROTTLE",
	AbsRUDDER:        "RUDDER",
	AbsWHEEL:         "WHEEL",
	AbsGAS:           "GAS",
	AbsBRAKE:         "BRAKE",
	AbsHAT0X:         "HAT0X",
	AbsHAT0Y:         "HAT0Y",
	AbsHAT1X:         "HAT1X",
	AbsHAT1Y:         "HAT1Y",
	AbsHAT2X:         "HAT2X",
	AbsHAT2Y:         "HAT2Y",
	AbsHAT3X:         "HAT3X",
	AbsHAT3Y:         "HAT3Y",
	AbsPRESSURE:      "PRESSURE",
	AbsDISTANCE:      "DISTANCE",
	AbsTILTX:         "TILT_X",
	AbsTILTY:         "TILT_Y",
	AbsTOOLWIDTH:     "TOOL_WIDTH",
	AbsVOLUME:        "VOLUME",
	AbsMISC:          "MISC",
	AbsRESERVED:      "RESERVED",
	AbsMTSLOT:        "MT_SLOT",
	AbsMTTOUCHMAJOR:  "MT_TOUCH_MAJOR",
	AbsMTTOUCHMINOR:  "MT_TOUCH_MINOR",
	AbsMTWIDTHMAJOR:  "MT_WIDTH_MAJOR",
	AbsMTWIDTHMINOR:  "MT_WIDTH_MINOR",
	AbsMTORIENTATION: "MT_ORIENTATION",
	AbsMTPOSITIONX:   "MT_POSITION_X",
	AbsMTPOSITIONY:   "MT_POSITION_Y",
	AbsMTTOOLTYPE:    "MT_TOOL_TYPE",
	AbsMTBLOBID:      "MT_BLOB_ID",
	AbsMTTRACKINGID:  "MT_TRACKING_ID",
	AbsMTPRESSURE:    "MT_PRESSURE",
	AbsMTDISTANCE:    "MT_DISTANCE",
	AbsMTTOOLX:       "MT_TOOL_X",
	AbsMTTOOLY:       "MT_TOOL_Y",
}

func (a Abs) String() string {
	str, found := absToString[a]
	if !found {
		return fmt.Sprintf("Unknown(%d)", a)
	}

	return str
}
