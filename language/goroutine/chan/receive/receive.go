package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	close(ch)

	// receive drained and chan
	if v, ok := <-ch; !ok {
		fmt.Println("closed and drained")
	} else {
		fmt.Println(v)
	}

	bufCh := make(chan int, 3)

	bufCh <- 1
	bufCh <- 2

	// after closed, cannot send, but can receive
	close(bufCh)

	for i := 0; i < 3; i++ {
		// receive drained and chan
		if v, ok := <-bufCh; !ok {
			fmt.Println("closed and drained")
		} else {
			fmt.Println(v)
		}
	}

}
