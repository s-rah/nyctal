package workspace

import (
	"fmt"
	"image"
	"image/color"

	"nyctal/model"
	"nyctal/utils"
	"sync"
)

type Surface struct {
	ID       int
	Parent   int
	Bounding image.Rectangle
	Client   model.Client
}

// A Workspace represents a collection of surfaces (windows)
type Workspace struct {
	uniq       int
	children   map[int]*Surface
	order      []int
	width      int
	height     int
	buffer     []byte
	lock       sync.Mutex
	focusChild int
	pointerPos image.Point
	keyboard   *model.Keyboard
	Quit       bool
	mx         float32
	my         float32
}

func NewWorkspace(width int, height int) model.Client {
	return &Workspace{uniq: 1, Quit: false, width: width, height: height, buffer: make([]byte, 4*width*height), children: make(map[int]*Surface), focusChild: -1, keyboard: model.NewKeyboardModel()}
}

func (ws *Workspace) Parent() model.Client {
	return nil // workspaces cannot be nested...for now
}

func (ws *Workspace) GetChild(id int) model.Client {
	ws.lock.Lock()
	defer ws.lock.Unlock()

	if c, ok := ws.children[id]; ok {
		return c.Client
	}
	fmt.Printf("[error] no child found %d\n", id)
	return nil
}

func (ws *Workspace) Resize(width int, height int) {
	utils.Debug("workspace", fmt.Sprintf("resize: %d %d", width, height))
	ws.lock.Lock()
	defer ws.lock.Unlock()
	ws.width = width
	ws.height = height
	ws.configureChildren()
}

func (ws *Workspace) configureChildren() {
	if len(ws.children) == 0 {
		return
	}
	if len(ws.children) == 1 {
		cidx := ws.order[0]
		if surface, ok := ws.children[cidx]; ok {
			surface.Client.Resize(ws.width, ws.height)
			surface.Bounding = image.Rect(0, 0, ws.width, ws.height)
		}
	} else {
		mainWidth := int(float64(ws.width) * 0.66666)
		ws.children[ws.order[0]].Client.Resize(mainWidth, ws.height)
		ws.children[ws.order[0]].Bounding = image.Rect(0, 0, mainWidth, ws.height)
		stackwidth := ws.width - mainWidth
		stackHeight := int(float64(ws.height) / float64(len(ws.children)-1))
		for i := 1; i < len(ws.order); i++ {
			ws.children[ws.order[i]].Client.Resize(stackwidth, stackHeight)
			ws.children[ws.order[i]].Bounding = image.Rect(mainWidth, stackHeight*(i-1), ws.width, stackHeight*i)
		}
	}
}

func (ws *Workspace) AddChild(id int, client model.Client) int {
	utils.Debug("workspace", fmt.Sprintf("adding child#%d", id))
	ws.lock.Lock()
	defer ws.lock.Unlock()

	cid := ws.uniq
	ws.children[cid] = &Surface{Client: client, Parent: id, ID: cid}
	ws.uniq += 1

	ws.order = append(ws.order, cid)
	if len(ws.children) == 1 {
		client.ProcessFocus()
	}
	ws.configureChildren()
	return cid
}

func (ws *Workspace) RemoveChild(id int) {
	utils.Debug("workspace", fmt.Sprintf("removing children#%d", id))
	ws.lock.Lock()
	defer ws.lock.Unlock()

	order := []int{}
	for i := 0; i < len(ws.order); i++ {
		if ws.order[i] != id {
			order = append(order, ws.order[i])
		}
	}

	ws.order = order
	delete(ws.children, id)
	ws.configureChildren()
}

func (ws *Workspace) RemoveChildren(id int) {
	utils.Debug("workspace", fmt.Sprintf("removing child#%d", id))
	ws.lock.Lock()
	defer ws.lock.Unlock()

	order := []int{}
	for i := 0; i < len(ws.order); i++ {
		if ws.children[ws.order[i]].Parent != id {
			order = append(order, ws.order[i])
		}
	}

	ws.order = order

	for _, child := range ws.children {
		if child.Parent == id {
			delete(ws.children, child.ID)
		}
	}

	ws.configureChildren()
}
func (ws *Workspace) Buffer() *model.BGRA {
	ws.lock.Lock()
	defer ws.lock.Unlock()

	img := model.EmptyBGRA(image.Rect(0, 0, ws.width, ws.height))

	for idx := range ws.children {
		ws.blitChild(idx, img)
	}

	if ws.focusChild == -1 {
		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				img.Set(ws.pointerPos.X+x, ws.pointerPos.Y+y, color.RGBA{R: 254})
			}
		}
	}
	return img
}

func (ws *Workspace) blitChild(idx int, img *model.BGRA) {
	//utils.Debug("workspace", fmt.Sprintf("blitting child %d", len(ws.children)))
	if len(ws.children) == 0 {
		return
	}
	mainSurface := ws.children[idx]
	buffer := mainSurface.Client.Buffer()

	if buffer == nil {
		utils.Debug("workspace", fmt.Sprintf("buffer for child %d is nil...", idx))
		mainSurface.Client.Resize(mainSurface.Bounding.Dx(), mainSurface.Bounding.Dy())
		return
	}

	if mainSurface.Bounding.Dx() != buffer.Bounds().Dx() || mainSurface.Bounding.Dy() != buffer.Bounds().Dy() {
		utils.Debug("workspace", fmt.Sprintf("buffer and surface bounds are not matched %v %v", mainSurface.Bounding, buffer.Bounds()))
		mainSurface.Client.Resize(mainSurface.Bounding.Dx(), mainSurface.Bounding.Dy())
	}
	model.DrawCopyOver(img, mainSurface.Bounding, buffer, image.Pt(0, 0))

}

func (ws *Workspace) AckFrame() {
	ws.lock.Lock()
	defer ws.lock.Unlock()
	for _, child := range ws.children {
		child.Client.AckFrame()
	}
}

func (ws *Workspace) ProcessKeyboardEvent(ev model.KeyboardEvent) {
	ws.lock.Lock()
	defer ws.lock.Unlock()
	ws.keyboard.ProcessKeyboardEvent(ev)

	downKeys := ws.keyboard.DownKeys()
	if downKeys[model.KB_ALT] && downKeys[model.KB_ESC] {
		ws.Quit = true
		return
	}

	if focusedChild, ok := ws.children[ws.focusChild]; ok {
		focusedChild.Client.ProcessKeyboardEvent(ev)
	}
}

func (ws *Workspace) ProcessPointerEvent(ev model.PointerEvent) bool {
	ws.lock.Lock()
	defer ws.lock.Unlock()

	if ev.Move != nil {
		ws.mx = ev.Move.MX
		ws.my = ev.Move.MY
		ws.pointerPos = image.Pt(int(ws.mx), int(ws.my))
	}

	if intersects, _, _ := utils.IntersectsRect(int(ws.mx), int(ws.my), image.Rect(0, 0, ws.width, ws.height)); !intersects {
		for id, surface := range ws.children {
			if id == ws.focusChild {
				surface.Client.HandlePointerLeave()
			} else {
				surface.Client.HandlePointerLeave()
			}
		}
	} else {
		for child, surface := range ws.children {
			if surface != nil {
				if intersects, lx, ly := utils.IntersectsRect(int(ws.mx), int(ws.my), surface.Bounding); intersects {
					if ev.Move != nil {
						ev.Move.MX = float32(lx)
						ev.Move.MY = float32(ly)
					}
					if surface.Client.ProcessPointerEvent(ev) {
						ws.focusChild = child
					}
				} else {
					surface.Client.HandlePointerLeave()
				}
			}
		}
	}
	return true

}

func (ws *Workspace) HandlePointerLeave() {
	ws.lock.Lock()
	defer ws.lock.Unlock()
	for id, surface := range ws.children {
		if id == ws.focusChild {
			surface.Client.HandlePointerLeave()
		} else {
			surface.Client.HandlePointerLeave()
		}
	}
}

func (ws *Workspace) ProcessFocus() {
	ws.lock.Lock()
	defer ws.lock.Unlock()
	if len(ws.order) >= 1 {
		ws.focusChild = ws.order[0]
	}
}

func (ws *Workspace) ProcessUnFocus() {
	ws.lock.Lock()
	defer ws.lock.Unlock()
	for _, surface := range ws.children {
		surface.Client.ProcessUnFocus()
	}
	ws.focusChild = -1
}
