package deque

import "errors"

var ErrEmpty = errors.New("Deque is empty")

type Deque[T any] struct {
	buf        []T
	head, tail int
	size       int
}

// NewDeque returns a pointer to a new, empty queue.
func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{
		buf: make([]T, InitialCapacity),
	}
}

const InitialCapacity = 8

func (d *Deque[T]) grow() {
	newBuf := make([]T, len(d.buf)*2)
	for i := 0; i < d.size; i++ {
		newBuf[i] = d.buf[(d.head+i)%len(d.buf)]
	}
	d.buf = newBuf
	d.head = 0
	d.tail = d.size
}

func (d *Deque[T]) PushFront(value T) {
	if d.size == len(d.buf) {
		d.grow()
	}
	d.head = (d.head - 1 + len(d.buf)) % len(d.buf)
	d.buf[d.head] = value
	d.size++
}

func (d *Deque[T]) PushBack(value T) {
	if d.size == len(d.buf) {
		d.grow()
	}
	d.buf[d.tail] = value
	d.tail = (d.tail + 1) % len(d.buf)
	d.size++
}

func (d *Deque[T]) PopFront() (T, bool) {
	if d.size == 0 {
		var zero T
		return zero, false
	}
	val := d.buf[d.head]
	d.head = (d.head + 1) % len(d.buf)
	d.size--
	return val, true
}

func (d *Deque[T]) PopBack() (T, bool) {
	if d.size == 0 {
		var zero T
		return zero, false
	}
	d.tail = (d.tail - 1 + len(d.buf)) % len(d.buf)
	val := d.buf[d.tail]
	d.size--
	return val, true
}

func (d *Deque[T]) Front() (T, bool) {
	if d.size == 0 {
		var zero T
		return zero, false
	}
	return d.buf[d.head], true
}

func (d *Deque[T]) Back() (T, bool) {
	if d.size == 0 {
		var zero T
		return zero, false
	}
	return d.buf[(d.tail-1+len(d.buf))%len(d.buf)], true
}

func (d *Deque[T]) Len() int {
	return d.size
}

func (d *Deque[T]) IsEmpty() bool {
	return d.size == 0
}

func (d *Deque[T]) Clear() {
	d.head = 0
	d.tail = 0
	d.size = 0
	d.buf = make([]T, InitialCapacity)
}
