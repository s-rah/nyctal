package evdev

import (
	"fmt"
	"os"
	"strings"
)

type devices []string

func (d *devices) hasDevice(str string) bool {
	for _, device := range *d {
		if strings.Contains(str, device) {
			return true
		}
	}

	return false
}

var mouseDevice = devices{"mouse"}
var restrictedDevices = devices{"mouse"}
var allowedDevices = devices{"keyboard", "logitech mx keys"}

// FindMouseDevice
func FindMouseDevice() string {
	path := "/sys/class/input/event%d/device/name"
	resolved := "/dev/input/event%d"

	for i := 0; i < 255; i++ {
		buff, err := os.ReadFile(fmt.Sprintf(path, i))

		// prevent from checking non-existant files
		if os.IsNotExist(err) {
			break
		}
		if err != nil {
			continue
		}

		deviceName := strings.ToLower(string(buff))

		if mouseDevice.hasDevice(deviceName) {
			return fmt.Sprintf(resolved, i)
		}
	}
	return ""
}

// Like FindKeyboardDevice, but finds all devices which contain keyword 'keyboard'
// Returns an array of file paths which contain keyboard events
func FindAllKeyboardDevices() []string {
	path := "/sys/class/input/event%d/device/name"
	resolved := "/dev/input/event%d"

	valid := make([]string, 0)

	for i := 0; i < 255; i++ {
		buff, err := os.ReadFile(fmt.Sprintf(path, i))

		// prevent from checking non-existant files
		if os.IsNotExist(err) {
			break
		}
		if err != nil {
			continue
		}

		deviceName := strings.ToLower(string(buff))

		if restrictedDevices.hasDevice(deviceName) {
			continue
		} else if allowedDevices.hasDevice(deviceName) {
			valid = append(valid, fmt.Sprintf(resolved, i))
		}
	}
	return valid
}
