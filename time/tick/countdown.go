package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	abort := make(chan struct{})
	go readInput(abort)

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)

		// select waits until some case is ready
		// if more than one cases are ready, it picks one randomly
		select {
		case <-abort:
			fmt.Printf("abort\n")
			// can generate non-zero exit code, but deferred funcs are not run
			os.Exit(1)
		case x := <-tick:
			fmt.Printf("%v\n", x)
		}
	}
}

func readInput(abort chan<- struct{}) {
	os.Stdin.Read(make([]byte, 1)) //read one byte
	abort <- struct{}{}
}
