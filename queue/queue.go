package queue

// Queue represents a generic FIFO (first-in, first-out) queue data structure.
// Elements are added to the back and removed from the front of the queue.
//
// To create a new queue:
//
//	var q goforge.Queue[int]
//	q := new(goforge.Queue[int])
//	q := &goforge.Queue[int]{}
//
// Example:
//
//	q.Push(42)
//	val, ok := q.Pop()
type Queue[T any] struct {
	data []T
}

// NewQueue returns a pointer to a new, empty queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Enqueue adds an item to the back of the queue.
func (q *Queue[T]) Enqueue(item T) {
	q.data = append(q.data, item)
}

// Dequeue removes and returns the front item of the queue.
// If the queue is empty, returns the zero value of T and false.
func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if len(q.data) == 0 {
		return zero, false
	}
	item := q.data[0]
	q.data = q.data[1:]
	return item, true
}

// Front returns the front item of the queue without removing it.
// If the queue is empty, returns the zero value of T and false.
func (q *Queue[T]) Front() (T, bool) {
	var zero T
	if q.IsEmpty() {
		return zero, false
	}
	return q.data[0], true
}

// Len returns the number of elements currently in the queue.
func (q *Queue[T]) Len() int {
	return len(q.data)
}

// IsEmpty returns true if the queue contains no elements.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

// Clear removes all elements from the queue. Capacity is retained.
func (q *Queue[T]) Clear() {
	q.data = make([]T, 0)
}
