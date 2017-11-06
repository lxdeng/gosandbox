package main

import (
	"fmt"
)

func main() {
	s := "abc"
	fmt.Printf("type of s is: %T, len of s is: %d\n", s, len(s))

    // strings are read-only slices of bytes with a bit of extra syntactic support from the language
	a := []byte(s)
	fmt.Printf("type of a is: %T\n", a)
	fmt.Println(a)

	for index, rune := range s {
		fmt.Printf("index=%d, rune=%v, type of rune is : %T\n", index, rune, rune)
	}

	// sub string
	s1 := s[1:2]
	fmt.Printf("s1 type is : %T\n", s1)
}
