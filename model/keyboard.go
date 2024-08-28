package model

const KB_CTRL = 29
const KB_SHIFT = 42
const KB_ALT = 56
const KB_TAB = 15
const KB_ESC = 1
const KB_SUPER = 125

// Keyboard maintains a very basic model of the state of the keyboard, goverened by KeyboardEvents
// Its main use is to track modifier and other key states for wl_keyboard events
// Secondly we use it in our compositor for checking compostor-level keybindings
type Keyboard struct {
	state map[int]bool // key positions
}

func NewKeyboardModel() *Keyboard {
	return &Keyboard{state: make(map[int]bool)}
}

func (k *Keyboard) DownKeys() map[int]bool {
	return k.state
}

func (k *Keyboard) ProcessKeyboardEvent(ev KeyboardEvent) {
	if ev.State == 0 {
		delete(k.state, int(ev.Key))
	} else {
		k.state[int(ev.Key)] = true
	}
}
