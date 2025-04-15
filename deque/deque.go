package deque

import "errors"

var ErrEmpty = errors.New("Deque is empty")

type Deque[T any] struct {
	Buf        []T
	head, tail int
	size       int
}

const InitialCapacity = 8

func (d *Deque[T]) grow() {
	newBuf := make([]T, len(d.Buf)*2)
	for i := 0; i < d.size; i++ {
		newBuf[i] = d.Buf[(d.head+i)%len(d.Buf)]
	}
	d.Buf = newBuf
	d.head = 0
	d.tail = d.size
}

func (d *Deque[T]) PushFront(value T) {
	if d.size == len(d.Buf) {
		d.grow()
	}
	d.head = (d.head - 1 + len(d.Buf)) % len(d.Buf)
	d.Buf[d.head] = value
	d.size++
}

func (d *Deque[T]) PushBack(value T) {
	if d.size == len(d.Buf) {
		d.grow()
	}
	d.Buf[d.tail] = value
	d.tail = (d.tail + 1) % len(d.Buf)
	d.size++
}

func (d *Deque[T]) PopFront() (T, bool) {
	if d.size == 0 {
		var zero T
		return zero, false
	}
	val := d.Buf[d.head]
	d.head = (d.head + 1) % len(d.Buf)
	d.size--
	return val, true
}

func (d *Deque[T]) PopBack() (T, bool) {
	if d.size == 0 {
		var zero T
		return zero, false
	}
	d.tail = (d.tail - 1 + len(d.Buf)) % len(d.Buf)
	val := d.Buf[d.tail]
	d.size--
	return val, true
}

func (d *Deque[T]) Front() (T, bool) {
	if d.size == 0 {
		var zero T
		return zero, false
	}
	return d.Buf[d.head], true
}

func (d *Deque[T]) Back() (T, bool) {
	if d.size == 0 {
		var zero T
		return zero, false
	}
	return d.Buf[(d.tail-1+len(d.Buf))%len(d.Buf)], true
}

func (d *Deque[T]) Len() int {
	return d.size
}

func (d *Deque[T]) Empty() bool {
	return d.size == 0
}

func (d *Deque[T]) Clear() {
	d.head = 0
	d.tail = 0
	d.size = 0
	d.Buf = make([]T, InitialCapacity)
}
