package main

import (
	"fmt"
	"runtime"
)

/*
 Observer how OS threads are used, by go scheduler.
 export GOMAXPROCS=1
 go run maxprocs.go
 export GOMAXPROCS=2
 go run maxprocs.go
*/

func main() {
	previous := runtime.GOMAXPROCS(4)
	fmt.Printf("GOMAXPROCS=%d\n", previous)

	for {
		go fmt.Print("0")
		fmt.Print("1")
	}
}
