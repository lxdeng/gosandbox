package main

import "fmt"

func counter(start int) (func() int, func()) {
	// if the value gets mutated, the same is reflected in closure
	ctr := func() int {
		return start
	}

	incr := func() {
		start++
	}

	// both ctr and incr have same reference to start
	// closures are created, but are not called
	return ctr, incr
}

func test2() {
	ctr, incr := counter(100) // All closures created together have same state.

    // the state are different for different creations of a closure
    // ctr1, incr1 are different from ctr, incr
	ctr1, incr1 := counter(100)

	fmt.Println("counter - ", ctr())
	fmt.Println("counter1 - ", ctr1())

	// incr by 1
	incr()
	fmt.Println("counter - ", ctr())
	fmt.Println("counter1 - ", ctr1())

	// incr1 by 2
	incr1()
	incr1()
	fmt.Println("counter - ", ctr())
	fmt.Println("counter1 - ", ctr1())
}
