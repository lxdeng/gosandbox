package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2))

	// pass in a slice
	s := []int{1, 2, 3}
	fmt.Println(sum(s...)) // ... is needed

	// nil slice
	s = []int(nil)
	fmt.Println(sum(s...))
}

// only the final parameter can have various numbers of arguments
func sum(nums ...int) int {
	// type is slice []int
	fmt.Printf("type=%T\n", nums)
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
