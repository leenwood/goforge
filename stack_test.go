package goforge

import (
	"testing"
)

func testPush[T comparable](t *testing.T, values []T) {
	stack := NewStack[T]()
	for _, v := range values {
		stack.Push(v)
	}
	if stack.Len() != len(values) {
		t.Errorf("Push: expected len %d, got %d", len(values), stack.Len())
	}
}

func testPop[T comparable](t *testing.T, values []T) {
	stack := NewStack[T]()
	for _, v := range values {
		stack.Push(v)
	}
	for i := len(values) - 1; i >= 0; i-- {
		val, ok := stack.Pop()
		if !ok || val != values[i] {
			t.Errorf("Pop: expected %v, got %v (ok=%v)", values[i], val, ok)
		}
	}
	_, ok := stack.Pop()
	if ok {
		t.Error("Pop: expected false on empty stack")
	}
}

func testPeek[T comparable](t *testing.T, value T) {
	stack := NewStack[T]()
	stack.Push(value)

	val, ok := stack.Peek()
	if !ok || val != value {
		t.Errorf("Peek: expected %v, got %v (ok=%v)", value, val, ok)
	}

	if stack.Len() != 1 {
		t.Errorf("Peek: expected stack to remain unchanged, got len %d", stack.Len())
	}

	stack.Pop()
	_, ok = stack.Peek()
	if ok {
		t.Error("Peek: expected false on empty stack")
	}
}

func testClear[T comparable](t *testing.T, values []T) {
	stack := NewStack[T]()
	for _, v := range values {
		stack.Push(v)
	}
	stack.Clear()
	if !stack.IsEmpty() {
		t.Error("Clear: expected stack to be empty")
	}
	if stack.Len() != 0 {
		t.Errorf("Clear: expected length 0, got %d", stack.Len())
	}
}

func testReverse[T comparable](t *testing.T, values []T) {
	stack := NewStack[T]()
	for _, v := range values {
		stack.Push(v)
	}
	stack.Reverse()
	for _, expected := range values {
		val, ok := stack.Pop()
		if !ok || val != expected {
			t.Errorf("Reverse: expected %v, got %v (ok=%v)", expected, val, ok)
		}
	}
}

func TestStack_Functions(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		testPush(t, []int{1, 2, 3})
		testPop(t, []int{1, 2, 3})
		testPeek(t, 42)
		testClear(t, []int{1, 2, 3, 4, 5})
		testReverse(t, []int{1, 2, 3})
	})

	t.Run("Float64", func(t *testing.T) {
		testPush(t, []float64{1.1, 2.2, 3.3})
		testPop(t, []float64{1.1, 2.2, 3.3})
		testPeek(t, 9.81)
		testClear(t, []float64{0.1, 0.2})
		testReverse(t, []float64{3.14, 2.71, 1.61})
	})

	t.Run("String", func(t *testing.T) {
		testPush(t, []string{"one", "two", "three"})
		testPop(t, []string{"one", "two", "three"})
		testPeek(t, "golang")
		testClear(t, []string{"a", "b", "c"})
		testReverse(t, []string{"first", "second", "third"})
	})

	t.Run("Rune", func(t *testing.T) {
		testPush(t, []rune{'a', 'b', 'c'})
		testPop(t, []rune{'a', 'b', 'c'})
		testPeek(t, 'x')
		testClear(t, []rune{'x', 'y', 'z'})
		testReverse(t, []rune{'1', '2', '3'})
	})
}
