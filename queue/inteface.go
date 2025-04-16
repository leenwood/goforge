package queue

// QueueInterface defines the behavior of a generic FIFO (first-in, first-out) queue.
//
// This interface is implemented by Queue[T] and can be used to abstract away
// the underlying queue implementation, making it easier to mock or substitute
// different implementations when needed.
type QueueInterface[T any] interface {
	// Push adds an item to the back of the queue.
	Push(item T)

	// Pop removes and returns the front item of the queue.
	// If the queue is empty, it returns the zero value of T and false.
	Pop() (T, bool)

	// Peek returns the front item of the queue without removing it.
	// If the queue is empty, it returns the zero value of T and false.
	Peek() (T, bool)

	// Len returns the number of elements currently in the queue.
	Len() int

	// IsEmpty returns true if the queue contains no elements.
	IsEmpty() bool

	// Clear removes all elements from the queue. Capacity is retained.
	Clear()
}
