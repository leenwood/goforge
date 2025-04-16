package queue

import "github.com/leenwood/goforge/base"

// QueueInterface defines the behavior of a generic FIFO (first-in, first-out) queue.
//
// It extends ContainerInterface and provides queue-specific operations such as
// Enqueue, Dequeue, and Front, allowing access to the front of the queue.
type QueueInterface[T any] interface {
	base.ContainerInterface[T]

	// Enqueue adds an item to the back of the queue.
	Enqueue(item T)

	// Dequeue removes and returns the front item of the queue.
	// If the queue is empty, it returns the zero value of T and false.
	Dequeue() (T, bool)

	// Front returns the front item of the queue without removing it.
	// If the queue is empty, it returns the zero value of T and false.
	Front() (T, bool)
}
