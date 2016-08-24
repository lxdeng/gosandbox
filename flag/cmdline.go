package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		count   = flag.Int("count", 1000, "total number of loops")
		debug   = flag.Bool("debug", false, "enable debug")
		version = flag.String("version", "1.0", "version string")
	)

	//level := ""
	//logLevel := &level
	logLevel := new(string) // the pointed string has default string zero value ""
	fmt.Printf("logLevel has value %v\n", *logLevel)

	flag.Parse()

	fmt.Printf("count has value %v\n", *count)
	fmt.Printf("debug has value %v\n", *debug)
	fmt.Printf("version has value %v\n", *version)
	fmt.Printf("logLevel has value %v\n", *logLevel)

	fmt.Println("\nPrintDefaults():")
	flag.PrintDefaults()
}
