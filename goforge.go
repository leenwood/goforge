package goforge

import (
	"github.com/leenwood/goforge/queue"
	"github.com/leenwood/goforge/stack"
)

func NewQueue[T any]() *queue.Queue[T] {
	return &queue.Queue[T]{}
}

func NewStack[T any]() *stack.Stack[T] {
	return &stack.Stack[T]{}
}
