package stack

import (
	"github.com/leenwood/goforge/base"
)

// StackInterface defines the behavior of a generic LIFO (last-in, first-out) stack.
//
// It extends ContainerInterface and provides stack-specific operations such as
// Push, Pop, and Peek, allowing direct manipulation of the top element.
type StackInterface[T any] interface {
	base.ContainerInterface[T]

	// Push adds an item to the top of the stack.
	Push(item T)

	// Pop removes and returns the top item of the stack.
	// If the stack is empty, it returns the zero value of T and false.
	Pop() (T, bool)

	// Peek returns the top item of the stack without removing it.
	// If the stack is empty, it returns the zero value of T and false.
	Peek() (T, bool)

	// Reverse reverses the order of elements in the stack in place.
	Reverse()
}
