package main

import (
	"fmt"
)

func outer(name string) {
	// variable in outer function
	text := "Modified " + name

	// foo is a inner function and has access to text variable, is a closure
	// closures have access to variables even after exiting this block
	foo := func() { // anonymous func
		fmt.Println(text)
	}

	// calling the closure
	foo()
}

func test1() {
	outer("hello")
}
