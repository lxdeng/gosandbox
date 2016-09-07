package main

import (
	"reflect"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}

	result := twoSum(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Error("expected", expected, "got", result)
	}
}
func TestTwoSum2(t *testing.T) {
    nums := []int{2, 7, 11, 15}
    target := 26 
    expected := []int{2, 3}

    result := twoSum(nums, target)
    if !reflect.DeepEqual(result, expected) {
        t.Error("expected", expected, "got", result)
    }
}
