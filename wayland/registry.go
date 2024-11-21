package wayland

import (
	"fmt"
	"sync"
	"sync/atomic"

	"nyctal/utils"
)

type Object interface {
	HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error
}

type NullObject struct {
}

func (no *NullObject) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {
	return fmt.Errorf("handle messsage called on nil object...this is not valid")
}

type Registry struct {
	objects   map[uint32]Object
	globalIdx atomic.Uint32
	lock      sync.Mutex
}

func NewRegistry() *Registry {
	return &Registry{objects: make(map[uint32]Object)}
}

func (r *Registry) Destroy(id uint32) {
	r.lock.Lock()
	defer r.lock.Unlock()
	delete(r.objects, id)
}

func (r *Registry) Get(id uint32) (Object, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if obj, ok := r.objects[id]; ok {
		return obj, nil
	} else {
		return nil, fmt.Errorf("unknown object")
	}
}

func (r *Registry) New(id uint32, obj Object) {
	r.lock.Lock()
	defer r.lock.Unlock()
	utils.Debug("registry", fmt.Sprintf("new object #%d = %T", id, obj))
	r.objects[id] = obj
}

func (r *Registry) FindSurfaces() []*Surface {
	r.lock.Lock()
	defer r.lock.Unlock()
	var surfaces []*Surface
	for _, object := range r.objects {
		if surface, ok := object.(*Surface); ok {
			surfaces = append(surfaces, surface)
		}
	}
	return surfaces

}

func (r *Registry) FindXDGSurfaces() []*XDG_Surface {
	r.lock.Lock()
	defer r.lock.Unlock()
	var surfaces []*XDG_Surface
	for _, object := range r.objects {
		if surface, ok := object.(*XDG_Surface); ok {
			surfaces = append(surfaces, surface)
		}
	}
	return surfaces
}

func (r *Registry) FindSeat() *Seat {
	r.lock.Lock()
	defer r.lock.Unlock()
	for _, object := range r.objects {
		if seat, ok := object.(*Seat); ok {
			return seat
		}
	}
	return nil
}

func (r *Registry) FindOutput() *Output {
	r.lock.Lock()
	defer r.lock.Unlock()
	for _, object := range r.objects {
		if seat, ok := object.(*Output); ok {
			return seat
		}
	}
	return nil
}
