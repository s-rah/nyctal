package wayland

import (
	"fmt"
	"sync"
)

type Object interface {
	HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error
	Destroy()
}

type NullObject struct {
	BaseObject
}

func (no *NullObject) HandleMessage(wsc *WaylandServerConn, packet *WaylandMessage) error {
	return fmt.Errorf("handle messsage called on nil object...this is not valid")
}

type BaseObject struct {
}

func (no *BaseObject) Destroy() {
}

type Registry struct {
	objects map[uint32]Object
	lock    sync.Mutex
}

func NewRegistry() *Registry {
	return &Registry{objects: make(map[uint32]Object)}
}

func (r *Registry) Close() {
	r.lock.Lock()
	defer r.lock.Unlock()
	for id, obj := range r.objects {
		obj.Destroy()
		delete(r.objects, id)
	}
}

func (r *Registry) Destroy(id uint32) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if obj, ok := r.objects[id]; ok {
		obj.Destroy()
	}
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
	if o, ok := r.objects[id]; ok {
		o.Destroy()
	}
	r.objects[id] = obj
}

func (r *Registry) FindDataDevice() *DataDevice {
	r.lock.Lock()
	defer r.lock.Unlock()
	for _, object := range r.objects {
		if seat, ok := object.(*DataDevice); ok {
			return seat
		}
	}
	return nil
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
