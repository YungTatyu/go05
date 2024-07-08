package test

import (
	"piscine"
	"reflect"
	"testing"
)

func expect(actual, expect []string, t *testing.T) {
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("expected %s, but got %s", expect, actual)
	}
}

func TestSplitWhiteSpaces(t *testing.T) {
	expect(piscine.SplitWhiteSpaces("this is\ntest\tyay"), []string{"this", "is", "test", "yay"}, t)
	expect(piscine.SplitWhiteSpaces("     this\t\ttest\n\n\n   "), []string{"this", "test"}, t)
	expect(piscine.SplitWhiteSpaces("     test\n\n\n   "), []string{"test"}, t)
	expect(piscine.SplitWhiteSpaces("     t e\ts\nt\n\n\n   "), []string{"t", "e", "s", "t"}, t)
	expect(piscine.SplitWhiteSpaces("\n\n   t"), []string{"t"}, t)
	expect(piscine.SplitWhiteSpaces("     \t\t\n\n\n   "), nil, t)
	expect(piscine.SplitWhiteSpaces(""), nil, t)
}
