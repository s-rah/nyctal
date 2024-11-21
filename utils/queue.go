package utils

type Queue[T any] struct {
	elements []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{nil}
}

func (queue *Queue[T]) Top() (T, bool) {
	var x T
	if len(queue.elements) > 0 {
		x = queue.elements[0]
		return x, true
	}
	return x, false
}

func (queue *Queue[T]) Push(key T) {
	queue.elements = append(queue.elements, key)
}

func (queue *Queue[T]) PushTop(key T) {
	queue.elements = append([]T{key}, queue.elements...)
}

func (queue *Queue[T]) Pop() (T, bool) {
	var x T
	if len(queue.elements) > 0 {
		x, queue.elements = queue.elements[0], queue.elements[1:]
		return x, true
	}
	return x, false
}

func (queue *Queue[T]) Empty() bool {
	return len(queue.elements) == 0
}

func (queue *Queue[T]) Inner() []T {
	return queue.elements
}
