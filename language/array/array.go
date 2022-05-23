package main

import (
	"fmt"
	"math"
)

func main() {
	b := [...]string{"Oregon", "Washington"}
	fmt.Printf("type of b is: %T, len of b is: %d\n", b, len(b))

	s := b[:]
	fmt.Printf("type of s is: %T, len of s is: %d\n", s, len(s))

	x, y := 1, 2

	if math.Abs(float64(x - y)) <= 3 {
		fmt.Println("true")
	}

}
