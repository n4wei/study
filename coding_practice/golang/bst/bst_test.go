package bst

import (
	"reflect"
	"testing"
	"strconv"
)

func assertEqual(t *testing.T, actual interface{}, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nActual:\n%v\nExpected:\n%v\n", actual, expected)
	}
}

func TestBST(t *testing.T) {
	bst := NewBST()
	bst.Insert("k5", "5")
	bst.Insert("k3", "3")
	bst.Insert("k2", "2")
	bst.Insert("k4", "4")
	bst.Insert("k0", "0")
	bst.Insert("k1", "1")
	bst.Insert("k7", "7")
	bst.Insert("k6", "6")
	bst.Insert("k9", "9")
	bst.Insert("k8", "8")
	assertEqual(t, bst.Len(), 10)

	iter := bst.Start()
	val := 0
	for !iter.Done() {
		assertEqual(t, iter.Value(), strconv.Itoa(val))
		iter.Next()
		val++
	}

	assertEqual(t, bst.Len(), 10)
	assertEqual(t, iter.Done(), true)
	assertEqual(t, iter.Value(), "")
	iter.Next()
	assertEqual(t, iter.Value(), "")
}

func TestLowerBound(t *testing.T) {
	bst := NewBST()
	bst.Insert("k5", "5")
	bst.Insert("k2", "2")
	bst.Insert("k0", "0")
	bst.Insert("k3", "3")
	bst.Insert("k7", "7")
	bst.Insert("k6", "6")
	bst.Insert("k8", "8")

	iter := bst.LowerBound("k3")
	assertEqual(t, iter.Value(), "3")
	iter.Next()
	assertEqual(t, iter.Value(), "5")
	iter.Next()
	assertEqual(t, iter.Value(), "6")

	iter = bst.LowerBound("k1")
	assertEqual(t, iter.Value(), "2")
	iter.Next()
	assertEqual(t, iter.Value(), "3")

	iter = bst.LowerBound("k0")
	assertEqual(t, iter.Value(), "0")
	iter.Next()
	assertEqual(t, iter.Value(), "2")

	iter = bst.LowerBound("k9")
	assertEqual(t, iter.Value(), "")
	assertEqual(t, iter.Done(), true)
}

func TestUpperBound(t *testing.T) {
	bst := NewBST()
	bst.Insert("k5", "5")
	bst.Insert("k2", "2")
	bst.Insert("k0", "0")
	bst.Insert("k3", "3")
	bst.Insert("k7", "7")
	bst.Insert("k6", "6")
	bst.Insert("k8", "8")

	iter := bst.UpperBound("k3")
	assertEqual(t, iter.Value(), "5")
	iter.Next()
	assertEqual(t, iter.Value(), "6")

	iter = bst.UpperBound("k1")
	assertEqual(t, iter.Value(), "2")
	iter.Next()
	assertEqual(t, iter.Value(), "3")

	iter = bst.UpperBound("k0")
	assertEqual(t, iter.Value(), "2")
	iter.Next()
	assertEqual(t, iter.Value(), "3")

	iter = bst.UpperBound("k8")
	assertEqual(t, iter.Value(), "")
	assertEqual(t, iter.Done(), true)
}

func TestFind(t *testing.T) {
	bst := NewBST()
	bst.Insert("k5", "5")
	bst.Insert("k2", "2")
	bst.Insert("k0", "0")
	bst.Insert("k3", "3")
	bst.Insert("k7", "7")
	bst.Insert("k6", "6")
	bst.Insert("k8", "8")

	iter := bst.Find("k3")
	assertEqual(t, iter.Value(), "3")
	iter.Next()
	assertEqual(t, iter.Value(), "5")

	iter = bst.Find("k1")
	assertEqual(t, iter.Value(), "")
	assertEqual(t, iter.Done(), true)

	iter = bst.Find("k")
	assertEqual(t, iter.Value(), "")
	assertEqual(t, iter.Done(), true)

	iter = bst.Find("k9")
	assertEqual(t, iter.Value(), "")
	assertEqual(t, iter.Done(), true)
}

func TestBSTRecursive(t *testing.T) {
	bst := NewBST()
	bst.InsertRecursive("k3", "3")
	bst.InsertRecursive("k2", "2")
	bst.InsertRecursive("k4", "4")
	bst.InsertRecursive("k1", "1")
	bst.InsertRecursive("k5", "5")
	bst.InsertRecursive("k0", "0")
	assertEqual(t, bst.Len(), 6)

	iter := bst.Start()
	val := 0
	for !iter.Done() {
		assertEqual(t, iter.Value(), strconv.Itoa(val))
		iter.Next()
		val++
	}
}