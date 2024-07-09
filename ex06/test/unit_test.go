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

func TestSplit(t *testing.T) {
	expect(piscine.Split(" this is test   ", " "), []string{"this", "is", "test"}, t)
	expect(piscine.Split("this is test isissi", "is"), []string{"th", " ", " test ", "si"}, t)
	expect(piscine.Split("sepsesepsepssepsepsepsepsesepse", "sep"), []string{"se", "s", "se", "se"}, t)
	expect(piscine.Split("sssssts s sss ", "s"), []string{"t", " ", " ", " "}, t)
	expect(piscine.Split(" this is test   ", ""), []string{" this is test   "}, t)
	expect(piscine.Split("", ""), nil, t)
	expect(piscine.Split("", " "), nil, t)
}
