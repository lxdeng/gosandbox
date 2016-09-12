package main

/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.
*/

import (
	"fmt"
)

func main() {
	test()
}

func isValid(s string) bool {
	bytes := []byte(s)
	stack := make([]byte, len(s))
	top := 0

	for _, b := range bytes {
		//fmt.Println(b)
		if b == '(' {
			stack[top] = ')'
			top++
		} else if b == '{' {
			stack[top] = '}'
			top++
		} else if b == '[' {
			stack[top] = ']'
			top++
		} else if top == 0 {
			return false
		} else {
			top--
			c := stack[top]
			if c != b {
				return false
			}
		}
	}

	return top == 0
}

func test() {
	result := isValid("[]{}")
	fmt.Println(result)

	result = isValid("[(])")
	fmt.Println(result)
}
