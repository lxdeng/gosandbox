package main

import (
	"fmt"
	"os"
	"strings"
)

// The first element of os.Args, os.Args[0], is the name of the command itself
func main() {
	printArgs1()
	printArgs2()
	printArgs3()
}

func printArgs1() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func printArgs2() {
	var s, sep string
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// efficient
func printArgs3() {
	fmt.Println(strings.Join(os.Args, " "))
}
