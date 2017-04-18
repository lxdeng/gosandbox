package main

import (
	"errors"
	"fmt"
	"io"
)

func main() {
	if _, err := f1(); err == io.EOF {
		fmt.Println("got io.EOF")
	} else {
		fmt.Println("did not get io.EOF")
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
