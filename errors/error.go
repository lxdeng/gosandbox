package main

import (
	"errors"
	"fmt"
	"io"
)

type User struct {
	Name string
}

func main() {

	testErrorChain()

	u1 := User{"Tom"}
	u2 := User{"Tom"}

	// Struct values are comparable if all their fields are comparable.
	// Two struct values are equal if their corresponding non-blank fields are equal.
	if u1 == u2 {
		fmt.Println("u1 == u2")
	} else {
		fmt.Println("u1 != u2")
	}

	//
	// io.EOF is of type pointer
	//
	if _, err := f1(); err == io.EOF {
		fmt.Println("got io.EOF")
		fmt.Printf("%T\n", io.EOF)
		fmt.Printf("%T\n", err)
	} else {
		fmt.Println("did not get io.EOF")
		fmt.Printf("%T\n", io.EOF)
		fmt.Printf("%T\n", err)
	}

	if _, err := f2(); err == io.EOF {
		fmt.Println("got io.EOF")
	} else {
		fmt.Println("did not get io.EOF")
	}
}

func f1() (int, error) {
	var myEOF = errors.New("EOF")
	return 0, myEOF
}

func f2() (int, error) {
	return 0, io.EOF
}

func testErrorChain() {
	r, err := checkPage()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}

func checkPage() (int, error) {
	_, err := checkParagraph()
	if err != nil {
        return 0, fmt.Errorf("page error: %v", err)
	} else {
        return 0, nil
	}
}

func checkParagraph() (int, error) {
	return 0, errors.New("paragraph error")
}
