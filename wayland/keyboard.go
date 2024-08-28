package wayland

import (
	"fmt"

	"nyctal/model"
	"nyctal/utils"
)

type Keyboard struct {
	id            uint32
	kb            *model.Keyboard
	activeSurface *Surface
	wsc           *WaylandServerConn
}

func NewKeyboard(id uint32, wsc *WaylandServerConn) *Keyboard {
	keyboard := &Keyboard{id: id, wsc: wsc, kb: model.NewKeyboardModel()}
	wsc.registry.New(id, keyboard)
	keyboard.SendKeyMap()
	return keyboard
}

// The only keymap specified in most wayland specs is the xkbcommon
// We are aiming for zero-dep, and the xkcommon format is very complicated and
// we have not (yet) written go-code to compile a real keymap
// as such we are always going to declare that we know nothing of keymaps, and instead
// default to sending raw scancodes.
// Note: Because of this we orix - ori running in an x11 window is dependent on a custom fork
// of minifb which surfaces scancodes instead of xkbcommon codes...
func (u *Keyboard) SendKeyMap() {
	pb := NewPacketBuilder(u.id, 0x00).
		WithUint(0).
		WithUint(0)

	utils.Debug(fmt.Sprintf("wl_keyboard#%d", u.id), fmt.Sprintf("keymap %d %d", 0, 0))
	u.wsc.SendMessageWithFd(pb.Build(), 0)
}

func (u *Keyboard) Enter(serial uint32, surface *Surface) {

	u.activeSurface = surface

	downKeys := u.kb.DownKeys()

	pb := NewPacketBuilder(u.id, 0x01).
		WithUint(serial).
		WithUint(u.activeSurface.id).
		WithUint(uint32(len(downKeys)))

	for key := range downKeys {
		pb.WithUint(uint32(key))

	}
	utils.Debug(fmt.Sprintf("wl_keyboard#%d", u.id), fmt.Sprintf("enter %d %d %x", serial, u.activeSurface.id, pb.Build()))

	u.wsc.SendMessage(pb.Build())
	u.sendModifiers(serial)
}

func (u *Keyboard) Leave(serial uint32) {

	if u.activeSurface != nil {
		pb := NewPacketBuilder(u.id, 0x02).
			WithUint(serial).
			WithUint(u.activeSurface.id)
		utils.Debug(fmt.Sprintf("wl_keyboard#%d", u.id), fmt.Sprintf("leave: %v", serial))
		u.wsc.SendMessage(pb.Build())
	}
}

func (u *Keyboard) ProcessKeyboardEvent(ev model.KeyboardEvent, serial uint32) {
	utils.Debug(fmt.Sprintf("wl_keyboard#%d", u.id), fmt.Sprintf("key: %v", ev))
	u.kb.ProcessKeyboardEvent(ev)
	if u.activeSurface != nil {
		utils.Debug("keyboard", "processing keyboard event")
		pb := NewPacketBuilder(u.id, 0x03).
			WithUint(serial).
			WithUint(ev.Time).
			WithUint(ev.Key).
			WithUint(ev.State)
		u.wsc.SendMessage(pb.Build())
		u.sendModifiers(serial)
	}
}

func (u *Keyboard) sendModifiers(serial uint32) {

	depressed := 0
	keys := u.kb.DownKeys()
	if keys[model.KB_SHIFT] {
		depressed |= 0x1
	}

	if keys[model.KB_CTRL] {
		depressed |= 0x4
	}

	pb := NewPacketBuilder(u.id, 0x04).
		WithUint(serial).
		WithUint(uint32(depressed)).
		WithUint(0).
		WithUint(0).
		WithUint(0)
	u.wsc.SendMessage(pb.Build())
}

func (u *Keyboard) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {

	switch packet.Opcode {
	case 0:
		// release
		wsc.registry.Destroy(u.id)
		return nil
	default:
		return fmt.Errorf("unknown opcode called on keyboard: %v", packet.Opcode)
	}

}
