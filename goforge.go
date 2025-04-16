package goforge

import (
	"github.com/leenwood/goforge/deque"
	"github.com/leenwood/goforge/queue"
	"github.com/leenwood/goforge/stack"
)

func NewQueue[T any]() queue.QueueInterface[T] {
	return &queue.Queue[T]{}
}

func NewStack[T any]() stack.StackInterface[T] {
	return &stack.Stack[T]{}
}

func NewDeque[T any]() deque.DequeInterface[T] {
	return &deque.Deque[T]{
		Buf: make([]T, deque.InitialCapacity),
	}
}
