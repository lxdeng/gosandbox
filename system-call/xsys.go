package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"log"
)

func main() {
	Getwd()

	Getuid()
}

func Getwd() {
	if wd, err := unix.Getwd(); err == nil {
		fmt.Printf("wd=%s\n", wd)
	} else {
		log.Fatal(err)
	}
}

func Getuid() {
	uid := unix.Getuid()
	fmt.Printf("uid=%d\n", uid)
}
