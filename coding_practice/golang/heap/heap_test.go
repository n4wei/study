package heap

import (
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, actual interface{}, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nActual:\n%v\nExpected:\n%v\n", actual, expected)
	}
}

func TestMinHeap(t *testing.T) {
	minFunc := func(a, b int) bool {
		if a < b {
			return true
		}
		return false
	}

	heap := NewHeap([]int{5,4,3,2,1}, minFunc)
	assertEqual(t, heap.Top(), 1)
	heap.Pop()
	assertEqual(t, heap.Top(), 2)
	heap.Pop()
	assertEqual(t, heap.Top(), 3)
	assertEqual(t, heap.Len(), 3)

	heap = NewHeap(nil, minFunc)
	heap.Push(5)
	heap.Push(4)
	heap.Push(3)
	heap.Push(2)
	heap.Push(1)
	assertEqual(t, heap.Len(), 5)
	assertEqual(t, heap.Top(), 1)
	heap.Pop()
	assertEqual(t, heap.Top(), 2)
	heap.Pop()
	assertEqual(t, heap.Top(), 3)
	heap.Pop()
	heap.Push(1)
	assertEqual(t, heap.Len(), 3)
	assertEqual(t, heap.Top(), 1)
	heap.Pop()
	assertEqual(t, heap.Top(), 4)
	heap.Pop()
	assertEqual(t, heap.Top(), 5)
	heap.Pop()
	assertEqual(t, heap.Top(), -1)
	assertEqual(t, heap.Len(), 0)
	heap.Pop()
	assertEqual(t, heap.Len(), 0)
}

func TestMaxHeap(t *testing.T) {
	maxFunc := func(a, b int) bool {
		if a > b {
			return true
		}
		return false
	}

	heap := NewHeap([]int{1,2,3,4,5}, maxFunc)
	assertEqual(t, heap.Top(), 5)
	heap.Pop()
	assertEqual(t, heap.Top(), 4)
	heap.Pop()
	assertEqual(t, heap.Top(), 3)
	assertEqual(t, heap.Len(), 3)

	heap = NewHeap(nil, maxFunc)
	heap.Push(1)
	heap.Push(2)
	heap.Push(3)
	heap.Push(4)
	heap.Push(5)
	assertEqual(t, heap.Len(), 5)
	assertEqual(t, heap.Top(), 5)
	heap.Pop()
	assertEqual(t, heap.Top(), 4)
	heap.Pop()
	assertEqual(t, heap.Top(), 3)
	heap.Pop()
	heap.Push(5)
	assertEqual(t, heap.Len(), 3)
	assertEqual(t, heap.Top(), 5)
	heap.Pop()
	assertEqual(t, heap.Top(), 2)
	heap.Pop()
	assertEqual(t, heap.Top(), 1)
	heap.Pop()
	assertEqual(t, heap.Top(), -1)
	assertEqual(t, heap.Len(), 0)
	heap.Pop()
	assertEqual(t, heap.Len(), 0)
}