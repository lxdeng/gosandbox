package main

import (
	"testing"
)

func TestKDiffPair1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	k := 1
	expected := 4

	result := findPairs(nums, k)
	if result != expected {
		t.Error("expected", expected, "got", result)
	}
}

func TestKDiffPair2(t *testing.T) {
	nums := []int{1, 3, 1, 5, 4}
	k := 0
	expected := 1

	result := findPairs(nums, k)
	if result != expected {
		t.Error("expected", expected, "got", result)
	}
}

func TestKDiffPair3(t *testing.T) {
	nums := []int{}
	k := 1
	expected := 0

	result := findPairs(nums, k)
	if result != expected {
		t.Error("expected", expected, "got", result)
	}
}

func TestKDiffPair4(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	k := -1
	expected := 0

	result := findPairs(nums, k)
	if result != expected {
		t.Error("expected", expected, "got", result)
	}
}
