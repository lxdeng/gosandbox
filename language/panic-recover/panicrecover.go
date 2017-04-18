package main

import (
	"log"
	"fmt"
)

func main() {
	func1()
}

func func1() {
	defer func() {
		if p := recover(); p != nil {
			msg := fmt.Errorf("internal error: %v, p type is %T ", p, p)
			log.Fatal(msg)
		}
	}()

	panic(99)
}
