package deque

import "github.com/leenwood/goforge/base"

// DequeInterface represents a double-ended queue, allowing efficient insertion
// and removal of elements from both the front and the back.
//
// It extends ContainerInterface to support basic inspection and control methods.
type DequeInterface[T any] interface {
	base.ContainerInterface[T]

	// PushFront adds an element to the front of the Deque.
	PushFront(value T)

	// PushBack adds an element to the back of the Deque.
	PushBack(value T)

	// PopFront removes and returns the element from the front of the Deque.
	// The second return value is false if the Deque is empty.
	PopFront() (T, bool)

	// PopBack removes and returns the element from the back of the Deque.
	// The second return value is false if the Deque is empty.
	PopBack() (T, bool)

	// Front returns the element at the front of the Deque without removing it.
	// The second return value is false if the Deque is empty.
	Front() (T, bool)

	// Back returns the element at the back of the Deque without removing it.
	// The second return value is false if the Deque is empty.
	Back() (T, bool)
}
