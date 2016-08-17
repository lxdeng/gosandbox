package main

import (
	"fmt"
)

func main() {
	testAppend()

	// built-in append() may use the original underlying array to store the appended,
	// or allocate a new underlying array, if the cap is not sufficient
	testAppend2()

	testNilSlice()
}

func testAppend() {
	s := make([]string, 0, 5)
	fmt.Printf("cap(s)=%d, len(s)=%d\n", cap(s), len(s))
	s1 := append(s, "a")
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println()
}

func testAppend2() {
	a := make([]int, 8, 8)
	for i := 0; i < 8; i++ {
		a[i] = i
	}

	Test(a)
	fmt.Println(a)
	fmt.Println()
}

func Test(slice []int) {
	fmt.Printf("cap(the slice)=%d\n", cap(slice))
	slice = append(slice, 100)
	fmt.Println(slice)
	fmt.Printf("cap(the new slice)=%d\n", cap(slice))
}

func testNilSlice() {
	//nilSlice := []int(nil)
	var nilSlice []int
	fmt.Printf("cap(nilSlice)=%d,  len(nilSlice)=%d\n", cap(nilSlice), len(nilSlice))

	//emptySlice := []int{}
	emptySlice := make([]int, 0)
	fmt.Printf("cap(emptySlice)=%d,  len(emptySlice)=%d\n", cap(emptySlice), len(emptySlice))

	emptySlice = append(emptySlice, 1)
	fmt.Printf("emptySlice=%v, empty slice can grow, but not nil slice", emptySlice)
	fmt.Println()
}
