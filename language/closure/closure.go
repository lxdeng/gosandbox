package main

/*
A closure is a function that you can pass around that retains the same “environment” as the one it had when it was created. In other words, the function defined in the closure ‘remembers’ the environment in which it was created.

*/

import (
	"fmt"
)

func main() {
	fmt.Println("test1")
	test1()

	fmt.Println("test2")
	test2()
}
