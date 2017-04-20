// Basic sends and receives on channels are blocking.
// However, we can use `select` with a `default` clause to
// implement _non-blocking_ sends, receives, and even
// non-blocking multi-way `select`s.

package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Here's a non-blocking receive. If a value is
	// available on `messages` then `select` will take
	// the `<-messages` `case` with that value. If not
	// it will immediately take the `default` case.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	//A send operation on an unbuffered channel blocks the sending goroutine
	//until another goroutine executes a corresponding receive on the same channel,
	//start a go routine to do receive, otherwise the following send
	//won't send since it's non-blocking send
	//uncomment the following to see the send going through
	/*
		go func() {
			<-messages
			fmt.Println("got message")
		}()
	*/

	time.Sleep(1 * time.Second)

	// A non-blocking send works similarly.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// We can use multiple `case`s above the `default`
	// clause to implement a multi-way non-blocking
	// select. Here we attempt non-blocking receives
	// on both `messages` and `signals`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
