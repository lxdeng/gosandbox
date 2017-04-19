package main

/*
Given an array of integers and an integer k, you need to find the number of unique k-diff pairs in the array. Here a k-diff pair is defined as an integer pair (i, j), where i and j are both numbers in the array and their absolute difference is k.

Example 1:

Input: [3, 1, 4, 1, 5], k = 2
Output: 2
Explanation: There are two 2-diff pairs in the array, (1, 3) and (3, 5).
Although we have two 1s in the input, we should only return the number of unique pairs.
*/

import (
	"fmt"
)

func main() {
	test()
}

func findPairs(nums []int, k int) int {
	set := make(map[int]bool)

	count := 0

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i]-nums[j] == k) || (nums[j]-nums[i] == k) {
				if !(set[nums[i]] && set[nums[j]]) {
					set[nums[i]] = true
					set[nums[j]] = true
					count++
					fmt.Printf("%d, %d\n", nums[i], nums[j])
				}
			}
		}
	}
	return count
}

func test() {
	nums := []int{3, 1, 4, 1, 5}
	result := findPairs(nums, 2)
	fmt.Println(result)
}
