package stringz

import (
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, actual interface{}, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nActual:\n%v\nExpected:\n%v\n", actual, expected)
	}
}

func TestIndex(t *testing.T) {
	assertEqual(t, Index("", ""), -1)
	assertEqual(t, Index("a", ""), -1)
	assertEqual(t, Index("", "b"), -1)
	assertEqual(t, Index("hello_world", "_"), 5)
	assertEqual(t, Index("hello_world", "llo"), 2)
	assertEqual(t, Index("hello_world", "x"), -1)
}

func TestNextIndex(t *testing.T) {
	assertEqual(t, NextIndex("", "", 0), -1)
	assertEqual(t, NextIndex("a", "", 0), -1)
	assertEqual(t, NextIndex("", "b", 0), -1)
	assertEqual(t, NextIndex("hello_world", "llo", 0), 2)
	assertEqual(t, NextIndex("hello_world", "llo", 3), -1)
	assertEqual(t, NextIndex("hello_world", "o", 5), 7)
}

func TestCompare(t *testing.T) {
	assertEqual(t, Compare("", ""), 0)
	assertEqual(t, Compare("a", ""), 1)
	assertEqual(t, Compare("", "b"), -1)
	assertEqual(t, Compare("abc", "abc"), 0)
	assertEqual(t, Compare("wor", "world"), -1)
	assertEqual(t, Compare("world", "wor"), 1)
}

func TestContains(t *testing.T) {
	assertEqual(t, Contains("", ""), false)
	assertEqual(t, Contains("a", ""), false)
	assertEqual(t, Contains("", "b"), false)
	assertEqual(t, Contains("hello_world", "hello"), true)
	assertEqual(t, Contains("hello_world", "_"), true)
	assertEqual(t, Contains("hello_world", "world"), true)
	assertEqual(t, Contains("hello_world", "x"), false)
}

func TestHasPrefix(t *testing.T) {
	assertEqual(t, HasPrefix("", ""), false)
	assertEqual(t, HasPrefix("a", ""), false)
	assertEqual(t, HasPrefix("", "b"), false)
	assertEqual(t, HasPrefix("hello_world", "hello"), true)
	assertEqual(t, HasPrefix("hello_world", "world"), false)
}

func TestHasSuffix(t *testing.T) {
	assertEqual(t, HasSuffix("", ""), false)
	assertEqual(t, HasSuffix("a", ""), false)
	assertEqual(t, HasSuffix("", "b"), false)
	assertEqual(t, HasSuffix("hello_world", "hello"), false)
	assertEqual(t, HasSuffix("hello_world", "world"), true)
}

func TestSplit(t *testing.T) {
	assertEqual(t, Split("", ""), []string{})
	assertEqual(t, Split("abc", ""), []string{"a","b","c"})
	assertEqual(t, Split("", "b"), []string{""})
	assertEqual(t, Split("hello_world", "*"), []string{"hello_world"})
	assertEqual(t, Split("hello_world", "_"), []string{"hello", "world"})
	assertEqual(t, Split("hello_world", "o_w"), []string{"hell", "orld"})
	assertEqual(t, Split("abcabca", "a"), []string{"", "bc", "bc", ""})
	assertEqual(t, Split("hello_world", "l"), []string{"he", "", "o_wor", "d"})
}

func TestReplace(t *testing.T) {
	assertEqual(t, Replace("", "", "", -1), "")
	assertEqual(t, Replace("", "b", "x", -1), "")
	assertEqual(t, Replace("abcabc", "b", "x", 1), "axcac")
	assertEqual(t, Replace("abcabc", "b", "x", -1), "axcaxc")
	assertEqual(t, Replace("abcabc", "", "x", -1), "xaxbxcxaxbx")
	assertEqual(t, Replace("abcabc", "b", "", -1), "acac")
	assertEqual(t, Replace("abcabca", "a", "x", -1), "xbcxbcx")
}

func TestJoin(t *testing.T) {
	assertEqual(t, Join(nil, ""), "")
	assertEqual(t, Join([]string{"a","b"}, ""), "ab")
	assertEqual(t, Join(nil, "x"), "")
	assertEqual(t, Join([]string{"a","b","c"}, "-"), "a-b-c")
}