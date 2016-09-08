package main

/*
Reverse digits of an integer.

Example1: x = 123, return 321
Example2: x = -123, return -321

Returns 0 when the reversed integer overflows.
*/

import (
	"fmt"
)

func main() {
	test()
}

func reverse(x int32) int32 {
	var r int64 = 0
	for x != 0 {
		r = r*10 + int64(x)%10
		x = x / 10
	}

	if int64(int32(r)) != r {
		return 0
	} else {
		return int32(r)
	}
}

func test() {
	result := reverse(-123)
	fmt.Println(result)

	result = reverse(1234567899)
	fmt.Println(result)
}
