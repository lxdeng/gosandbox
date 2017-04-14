package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("./test.txt")
	check(err)

	// Scanner Scan/Text seems to get a line each time, instead of a token
	input := bufio.NewScanner(f)
	for input.Scan() {
		fmt.Printf("Scan() returns true, text: %s\n", input.Text())
	}

	if input.Err() == nil {
		fmt.Printf("Scanner reaches EOF\n")
	}

}
