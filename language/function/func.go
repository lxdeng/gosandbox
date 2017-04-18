package main

import (
	"fmt"
)

func main() {
	// Functions are first-class values in Go: like other values, function values have types,
	//  and they may be assigned to variables or passed to or returned from functions.
	f := square
	fmt.Printf("f's type is %T\n", f)
}

func square(x int) int {
	return x * x
}
