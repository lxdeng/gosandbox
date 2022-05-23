package main

import (
  "fmt"
)

type Number interface {
	Increase()
}

type MyInt int32

func (x MyInt) Increase() {
	x += 1
}

func main() {
	var my MyInt = 10

	var num Number = my

	num.Increase()

	fmt.Println(my)
	fmt.Println(num)
}