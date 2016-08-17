package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"log"
)

func main() {
	Getwd()

	Getuid()

	Environ()
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

func Environ() {
	envs := unix.Environ()
	for _, e := range envs {
		fmt.Printf("%s\n", e)
	}
}
