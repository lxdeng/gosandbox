package main

import (
	"fmt"
)

func main() {
	b := [...]string{"Oregon", "Washington"}
	fmt.Printf("type of b is: %T, len of b is: %d\n", b, len(b))

	s := b[:]
	fmt.Printf("type of s is: %T, len of s is: %d\n", s, len(s))
}
