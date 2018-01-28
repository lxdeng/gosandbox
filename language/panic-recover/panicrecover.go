package main

import (
	"log"
	"fmt"
)

func main() {
	func1()

	fmt.Println("resume normal executing")
}

func func1() {
	defer func() {
		if p := recover(); p != nil {
			msg := fmt.Errorf("internal error: %v, p type is %T ", p, p)
			log.Println("recovered from panicking: ", msg)
		}
	}()

	panic(99)
}
