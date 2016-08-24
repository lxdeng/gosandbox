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
		for i := 0; ; i++ {
			naturals <- i
			time.Sleep(3 * time.Second)
		}
	}()

	// Squarer
	go func() {
		for {
			i := <-naturals
			squares <- i * i
		}
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}
