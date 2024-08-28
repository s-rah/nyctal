package utils

type Stack[T any] struct {
	elements []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

func (stack *Stack[T]) Push(key T) {
	stack.elements = append(stack.elements, key)
}

func (stack *Stack[T]) Top() (T, bool) {
	var x T
	if len(stack.elements) > 0 {
		x = stack.elements[len(stack.elements)-1]
		return x, true
	}
	return x, false
}

func (stack *Stack[T]) Pop() (T, bool) {
	var x T
	if len(stack.elements) > 0 {
		x, stack.elements = stack.elements[len(stack.elements)-1], stack.elements[:len(stack.elements)-1]
		return x, true
	}
	return x, false
}

func (stack *Stack[T]) Empty() bool {
	return len(stack.elements) == 0
}

func (stack *Stack[T]) Inner() []T {
	return stack.elements
}
