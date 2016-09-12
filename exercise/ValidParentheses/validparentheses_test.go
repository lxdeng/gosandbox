package main

import (
	"testing"
)

func TestValid1(t *testing.T) {
	s := "[]{"
	result := isValid(s)
	expected := false
	if expected != result {
		t.Error("expected", expected, "got", result)
	}
}

func TestValid2(t *testing.T) {
	s := "{[]()}"
    result := isValid(s)
    expected := true
    if expected != result {
        t.Error("expected", expected, "got", result)
    }
}
