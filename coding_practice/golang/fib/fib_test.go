package fib

import (
	"reflect"
	"testing"
)

var expected []int = []int{0,1,1,2,3,5,8,13,21,34}

func assertEqual(t *testing.T, actual interface{}, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nActual:\n%v\nExpected:\n%v\n", actual, expected)
	}
}

func TestFibRecursive(t *testing.T) {
	actual := []int{}
	for i:=0; i<10; i++ {
		actual = append(actual, fibRecursive(i))
	}
	assertEqual(t, actual, expected)
}

func TestFibRecursiveMemo(t *testing.T) {
	actual := []int{}
	for i:=0; i<10; i++ {
		actual = append(actual, fibRecursiveMemo(i))
	}
	assertEqual(t, actual, expected)
}

func TestFibIterative(t *testing.T) {
	actual := []int{}
	for i:=0; i<10; i++ {
		actual = append(actual, fibIterative(i))
	}
	assertEqual(t, actual, expected)
}