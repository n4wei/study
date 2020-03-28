package stack

import (
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, actual interface{}, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nActual:\n%v\nExpected:\n%v\n", actual, expected)
	}
}

func TestStack(t *testing.T) {
	strStack := NewStringStack()
	assertEqual(t, strStack.Pop(), "")
	strStack.Push("a")
	strStack.Push("b")
	strStack.Push("c")
	assertEqual(t, strStack.Len(), 3)
	assertEqual(t, strStack.Pop(), "c")
	assertEqual(t, strStack.Pop(), "b")
	assertEqual(t, strStack.Len(), 1)
	strStack.Push("d")
	assertEqual(t, strStack.Len(), 2)
	assertEqual(t, strStack.Pop(), "d")
	assertEqual(t, strStack.Pop(), "a")
	assertEqual(t, strStack.Len(), 0)
}