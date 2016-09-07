package main

import (
	"reflect"
	"testing"
)

func TestAddTwo1(t *testing.T) {
	nums1 := []int{2, 4, 3}
	nums2 := []int{5, 6, 4}
	expected := []int{7, 0, 8}

	l1 := makeList(nums1)
	l2 := makeList(nums2)

	result := addTwoNumbers(l1, l2)
	resultList := makeSlice(result)

	if !reflect.DeepEqual(resultList, expected) {
		t.Error("expected", expected, "got", resultList)
	}
}

func TestTwoSum2(t *testing.T) {
    nums1 := []int{2, 4, 3}
    nums2 := []int{8}
    expected := []int{0, 5, 3}

    l1 := makeList(nums1)
    l2 := makeList(nums2)

    result := addTwoNumbers(l1, l2)
    resultList := makeSlice(result)

    if !reflect.DeepEqual(resultList, expected) {
        t.Error("expected", expected, "got", resultList)
    }
}
