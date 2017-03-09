package main

import (
	"fmt"
	"github.com/lxdeng/gosandbox/language/libstringutil"
)

func main() {
	fmt.Printf("func1=%d\n", func1(5))

    /* the imported "github.com/lxdeng/gosandbox/language/libstringutil" package
     is referred to using the name "stringutil" */
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
	fmt.Println()
}

func func1(x int) int {
	return x + 1
}
