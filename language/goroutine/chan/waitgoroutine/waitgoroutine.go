package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("enter main")

	done := make(chan struct{})
	go child(done)
	<-done
	fmt.Println("child done")

	var done2 sync.WaitGroup
	done2.Add(1)
	go child2(&done2)
	done2.Wait()
	fmt.Println("main exit")
}

func child(done chan<- struct{}) {
	fmt.Println("enter child")
	done <- struct{}{}
	fmt.Println("enter child")
}

func child2(done *sync.WaitGroup) {
	fmt.Println("enter child2")
	done.Done()
	fmt.Println("enter child2")
}
