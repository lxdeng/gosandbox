package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("start to count down")

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		fmt.Println("read a char")
		abort <- struct{}{}
	}()

	for i := 20; i > 0; i-- {
		println(i)
		select {
		case <-time.Tick(1 * time.Second):
			//
		case <-abort:
			log.Fatal("aborted")
		}
	}

	fmt.Println("lauch")
}
