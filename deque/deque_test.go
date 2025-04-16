package deque

import (
	"testing"
)

func testPushPopFrontBack[T comparable](t *testing.T, values []T) {
	q := NewDeque[T]()

	// PushBack: 1, 2, 3
	for _, v := range values {
		q.PushBack(v)
	}
	// PushFront: 3, 2, 1 → на переднюю часть дека
	for i := len(values) - 1; i >= 0; i-- {
		q.PushFront(values[i])
	}

	expectedLen := 2 * len(values)
	if q.Len() != expectedLen {
		t.Errorf("Expected len %d, got %d", expectedLen, q.Len())
	}

	// Первыми выйдут значения, добавленные через PushFront (в прямом порядке)
	for i := 0; i < len(values); i++ {
		val, ok := q.PopFront()
		if !ok || val != values[i] {
			t.Errorf("PopFront: expected %v, got %v (ok=%v)", values[i], val, ok)
		}
	}

	// Затем значения, добавленные через PushBack — надо доставать в обратном порядке
	for i := len(values) - 1; i >= 0; i-- {
		val, ok := q.PopBack()
		if !ok || val != values[i] {
			t.Errorf("PopBack: expected %v, got %v (ok=%v)", values[i], val, ok)
		}
	}

	if !q.IsEmpty() {
		t.Error("Expected deque to be empty after all pops")
	}
	if q.Len() != 0 {
		t.Errorf("Expected length 0, got %d", q.Len())
	}
}

func testPeekFrontBack[T comparable](t *testing.T, first, last T) {
	q := NewDeque[T]()
	q.PushFront(first)
	q.PushBack(last)

	front, ok := q.Front()
	if !ok || front != first {
		t.Errorf("Front: expected %v, got %v (ok=%v)", first, front, ok)
	}

	back, ok := q.Back()
	if !ok || back != last {
		t.Errorf("Back: expected %v, got %v (ok=%v)", last, back, ok)
	}

	if q.Len() != 2 {
		t.Errorf("Expected len 2, got %d", q.Len())
	}

	q.PopFront()
	q.PopBack()

	if _, ok = q.Front(); ok {
		t.Error("Expected Front to return false on empty deque")
	}
	if _, ok = q.Back(); ok {
		t.Error("Expected Back to return false on empty deque")
	}
}

func testClear[T comparable](t *testing.T, values []T) {
	q := NewDeque[T]()
	for _, v := range values {
		q.PushBack(v)
	}
	q.Clear()

	if !q.IsEmpty() {
		t.Error("Clear: expected deque to be empty")
	}
	if q.Len() != 0 {
		t.Errorf("Clear: expected length 0, got %d", q.Len())
	}

	if _, ok := q.PopFront(); ok {
		t.Error("PopFront: expected false after Clear")
	}
	if _, ok := q.PopBack(); ok {
		t.Error("PopBack: expected false after Clear")
	}
	if _, ok := q.Front(); ok {
		t.Error("Front: expected false after Clear")
	}
	if _, ok := q.Back(); ok {
		t.Error("Back: expected false after Clear")
	}
}

func TestDeque_Functions(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		testPushPopFrontBack(t, []int{1, 2, 3})
		testPeekFrontBack(t, 10, 20)
		testClear(t, []int{100, 200, 300})
	})

	t.Run("Float64", func(t *testing.T) {
		testPushPopFrontBack(t, []float64{1.1, 2.2, 3.3})
		testPeekFrontBack(t, 4.4, 5.5)
		testClear(t, []float64{6.6, 7.7})
	})

	t.Run("String", func(t *testing.T) {
		testPushPopFrontBack(t, []string{"one", "two", "three"})
		testPeekFrontBack(t, "start", "end")
		testClear(t, []string{"a", "b", "c"})
	})

	t.Run("Rune", func(t *testing.T) {
		testPushPopFrontBack(t, []rune{'x', 'y', 'z'})
		testPeekFrontBack(t, 'α', 'ω')
		testClear(t, []rune{'1', '2', '3'})
	})
}
