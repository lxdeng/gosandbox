package main

import (
	"fmt"
)

func main() {
	fmt.Printf("func1=%d\n", func1(5))
}

func func1(x int) int {
	return x + 1
}
