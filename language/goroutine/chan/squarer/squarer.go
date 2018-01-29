package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for i := 0; i < 10; i++ {
			// uni-directional channel
			var naturals2 chan<- int
			naturals2 = naturals
			naturals2 <- i
			time.Sleep(1 * time.Second)
		}
		// closing a channel to tell receiving end
		close(naturals)
	}()

	// Squarer
	go func() {
		//
		// receive on a closed channel gets a zero value, if not using range
		// x, ok := <-naturals get ok==false, if channel closed
		//
		for i := range naturals {
			squares <- i * i
		}
		close(squares)

	}()

	// Printer (in main goroutine)
	for i := range squares {
		fmt.Println(i)
	}

	fmt.Println("drained all")
}
