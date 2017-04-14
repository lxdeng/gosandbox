package main

import (
	"fmt"
)

func main() {
	testAppend()

	// built-in append() may use the original underlying array to store the appended,
	// or allocate a new underlying array, if the cap is not sufficient
	testAppend2()

	testAppendAnotherSlice()

	testAppendStringToByteSlice()

	testNilSlice()
}

func testAppendAnotherSlice() {
	fmt.Println("testAppendAnotherSlice")
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s3 := append(s1, s2...)
	fmt.Println(s3)
	fmt.Println()
}

func testAppend() {
	fmt.Println("testAppend")
	s := make([]string, 0, 5)
	fmt.Printf("cap(s)=%d, len(s)=%d\n", cap(s), len(s))
	s1 := append(s, "a")
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println()
}

func testAppend2() {
	fmt.Println("testAppend2")
	a := make([]int, 8, 8)
	for i := 0; i < 8; i++ {
		a[i] = i
	}

	fmt.Println("before calling Test on the slice:", a)

	Test(a)

	fmt.Println("after calling Test on the slice:", a)

	fmt.Println(a)

	TestPointer(&a)
	fmt.Println("after calling TestPointer on the slice:", a)
	fmt.Println()
}

// slice is a reference to array, efficient for passing
func Test(slice []int) {
	fmt.Printf("cap(the slice)=%d\n", cap(slice))
	slice = append(slice, 100)

	fmt.Println("inside Test:", slice)
	fmt.Printf("cap(the slice)=%d\n", cap(slice))
}

// using pointer to change the slice(the reference itself)
func TestPointer(pslice *[]int) {
	*pslice = append(*pslice, 101)
}

// As a special case, it is legal to append a string to a byte slice
func testAppendStringToByteSlice() {
	fmt.Println("testAppendStringToByteSlice")
	s := []byte("hello")
	s = append(s, " world"...)
	fmt.Println(s)
	fmt.Println()
}

func testNilSlice() {
	fmt.Println("testNilSlice")
	nilSlice := []int(nil)
	//var nilSlice []int
	fmt.Printf("cap(nilSlice)=%d,  len(nilSlice)=%d\n", cap(nilSlice), len(nilSlice))
	nilSlice = append(nilSlice, 99)
	fmt.Printf("aftern append nil slice: cap(nilSlice)=%d,  len(nilSlice)=%d\n", cap(nilSlice), len(nilSlice))

	//emptySlice := []int{}
	emptySlice := make([]int, 0)
	fmt.Printf("cap(emptySlice)=%d,  len(emptySlice)=%d\n", cap(emptySlice), len(emptySlice))

	emptySlice = append(emptySlice, 1)
	fmt.Printf("emptySlice=%v, empty slice can grow, nil slice can grow also", emptySlice)
	fmt.Println()
}
