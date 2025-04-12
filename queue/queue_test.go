package queue

import (
	"testing"
)

func testPush[T comparable](t *testing.T, values []T) {
	q := NewQueue[T]()
	for _, v := range values {
		q.Push(v)
	}
	if q.Len() != len(values) {
		t.Errorf("Push: expected len %d, got %d", len(values), q.Len())
	}
}

func testPop[T comparable](t *testing.T, values []T) {
	q := NewQueue[T]()
	for _, v := range values {
		q.Push(v)
	}
	for _, expected := range values {
		val, ok := q.Pop()
		if !ok || val != expected {
			t.Errorf("Pop: expected %v, got %v (ok=%v)", expected, val, ok)
		}
	}
	_, ok := q.Pop()
	if ok {
		t.Error("Pop: expected false on empty queue")
	}
}

func testPeek[T comparable](t *testing.T, value T) {
	q := NewQueue[T]()
	q.Push(value)

	val, ok := q.Peek()
	if !ok || val != value {
		t.Errorf("Peek: expected %v, got %v (ok=%v)", value, val, ok)
	}

	if q.Len() != 1 {
		t.Errorf("Peek: expected queue to remain unchanged, got len %d", q.Len())
	}

	q.Pop()
	_, ok = q.Peek()
	if ok {
		t.Error("Peek: expected false on empty queue")
	}
}

func testClear[T comparable](t *testing.T, values []T) {
	q := NewQueue[T]()
	for _, v := range values {
		q.Push(v)
	}
	q.Clear()
	if !q.IsEmpty() {
		t.Error("Clear: expected queue to be empty")
	}
	if q.Len() != 0 {
		t.Errorf("Clear: expected length 0, got %d", q.Len())
	}
}

func TestQueue_Functions(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		testPush(t, []int{1, 2, 3})
		testPop(t, []int{1, 2, 3})
		testPeek(t, 42)
		testClear(t, []int{1, 2, 3, 4, 5})
	})

	t.Run("Float64", func(t *testing.T) {
		testPush(t, []float64{1.1, 2.2, 3.3})
		testPop(t, []float64{1.1, 2.2, 3.3})
		testPeek(t, 9.81)
		testClear(t, []float64{0.1, 0.2})
	})

	t.Run("String", func(t *testing.T) {
		testPush(t, []string{"one", "two", "three"})
		testPop(t, []string{"one", "two", "three"})
		testPeek(t, "golang")
		testClear(t, []string{"a", "b", "c"})
	})

	t.Run("Rune", func(t *testing.T) {
		testPush(t, []rune{'a', 'b', 'c'})
		testPop(t, []rune{'a', 'b', 'c'})
		testPeek(t, 'x')
		testClear(t, []rune{'x', 'y', 'z'})
	})
}
