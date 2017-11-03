package main

/*
// Everything in comments above the import "C" is C code and will be compiles with the GCC.
// Make sure you have a GCC installed.

int addInC(int a, int b) {
    return a + b;
}
*/
import "C"

import "fmt"

func main() {
	a := 3
	b := 5

	var sum C.int = C.addInC(C.int(a), C.int(b))

	fmt.Println("Add in C:", a, "+", b, "=", int(sum))
}
