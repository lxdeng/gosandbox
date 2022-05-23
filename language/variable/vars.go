package main

import (
	"fmt"
)

// g1 is of type int. Type int is system dependent, and its size depends on the computer architecture.
var g1 = 1

// short form not allowed for package global var
//g2 := 1.0

// constants are a var whose value (of certain type) cannot be changed
const G = "abc"

func main() {
	g2 := 1.0

	fmt.Printf("g1's type is %T\n", g1)
	fmt.Printf("g2's type is %T\n", g2)
	fmt.Printf("G's type is %T\n", G)
}
