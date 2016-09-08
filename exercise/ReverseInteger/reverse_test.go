package main

import (
	"testing"
)

func TestReverse1(t *testing.T) {
	var x int32 = -123
	var expected int32 = -321

	result := reverse(x)
	if expected != result {
		t.Error("expected", expected, "got", result)
	}
}

func TestReverse2(t *testing.T) {
	var x int32 = 1234567899
    var expected int32 = 0

    result := reverse(x)
    if expected != result {
        t.Error("expected", expected, "got", result)
    }
}
