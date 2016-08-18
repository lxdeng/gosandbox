package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Print(err)
		return nil
	}
	fmt.Printf("%s isDir=%v\n", path, info.IsDir())
	return nil
}

func main() {
	log.SetFlags(log.Lshortfile)
	dir := os.Args[1]
	err := filepath.Walk(dir, visit)
	if err != nil {
		log.Fatal(err)
	}
}
