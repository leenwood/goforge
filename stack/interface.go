package stack

// StackInterface defines the behavior of a generic LIFO (last-in, first-out) stack.
//
// This interface is implemented by Stack[T] and can be used to abstract away
// the underlying stack implementation, allowing easier testing and replacement.
type StackInterface[T any] interface {
	// Push adds an item to the top of the stack.
	Push(item T)

	// Pop removes and returns the top item of the stack.
	// If the stack is empty, it returns the zero value of T and false.
	Pop() (T, bool)

	// Peek returns the top item of the stack without removing it.
	// If the stack is empty, it returns the zero value of T and false.
	Peek() (T, bool)

	// Len returns the number of elements currently in the stack.
	Len() int

	// IsEmpty returns true if the stack contains no elements.
	IsEmpty() bool

	// Clear removes all elements from the stack. Capacity is retained.
	Clear()

	// Reverse reverses the order of elements in the stack in place.
	Reverse()
}
