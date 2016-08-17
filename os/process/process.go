package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	getInitProcess()
}

func getInitProcess() {
	// even if an invalid process id, like -1, is given on Linux, it does not return an error
	p, err := os.FindProcess(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("pid=%d\n", p.Pid)
}
