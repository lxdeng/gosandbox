package main

import (
	"fmt"
	"sync"
)

type student struct {
	Name string
	Age  int
}

func main() {
	s := student{"Tom", 20}
	fmt.Println(s)
	fmt.Printf("s' address=%p\n", &s)

	schan := make(chan student)

	var done sync.WaitGroup
	done.Add(1)

	go func() {
		defer done.Done()
		s1 := <-schan
		fmt.Println(s1)
		s1.Age = 21
	}()

	schan <- s // send statement sends a value to a channel

	done.Wait()
	fmt.Println(s)
}
