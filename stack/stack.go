package stack

// Stack represents a generic LIFO (last-in, first-out) stack data structure.
// Elements are added to and removed from the top of the stack.
//
// To create a new stack:
//
//	var s goforge.Stack[int]
//	s := new(goforge.Stack[int])
//	s := &goforge.Stack[int]{}
//
// Example:
//
//	s.Push(42)
//	val, ok := s.Pop()
type Stack[T any] struct {
	data []T
}

// NewStack returns a pointer to a new, empty stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds an item to the top of the stack.
func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

// Pop removes and returns the top item of the stack.
// If the stack is empty, returns the zero value of T and false.
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.data) == 0 {
		return zero, false
	}
	last := len(s.data) - 1
	item := s.data[last]
	s.data = s.data[:last]
	return item, true
}

// Peek returns the top item of the stack without removing it.
// If the stack is empty, returns the zero value of T and false.
func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if s.IsEmpty() {
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

// Len returns the number of elements currently in the stack.
func (s *Stack[T]) Len() int {
	return len(s.data)
}

// IsEmpty returns true if the stack contains no elements.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Clear removes all elements from the stack. Capacity is retained.
func (s *Stack[T]) Clear() {
	s.data = make([]T, 0)
}

// Reverse reverses the order of elements in the stack in place.
func (s *Stack[T]) Reverse() {
	for i, j := 0, len(s.data)-1; i < j; i, j = i+1, j-1 {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
}
