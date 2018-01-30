package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var (
    mu      sync.Mutex // guards balance
    done = false
)

func main() {
	go checkInput()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go child(i, &wg)
	}

	wg.Wait()
	fmt.Println("main exit")
}

func child(id int, wg *sync.WaitGroup) {
	fmt.Printf("enter child %d\n", id)

	for i := 0; i < 5; i++ {
		if cancelled() {
			break
		}
		fmt.Printf("child %d now sleep %d\n", id, i+1)
		time.Sleep(10 * time.Second)
	}

	wg.Done()
	fmt.Printf("child %d exit\n", id)
}

func checkInput() {
	os.Stdin.Read(make([]byte, 1))

	mu.Lock()
	done = true
	mu.Unlock()
}

// utilize a closed-chan return zero value of channel type
func cancelled() bool {
	mu.Lock()
	defer mu.Unlock()
	return done;
}
