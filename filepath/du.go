package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir, sorted by name
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
    defer func() { <-sema }() // release token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, fileSizes chan<- int64, n *sync.WaitGroup) {
	defer n.Done()

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			//fileSizes <- entry.Size() //ignore directories
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, fileSizes, n)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args() // args without the command itself
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fSizes := make(chan int64)

	var n sync.WaitGroup

	for _, root := range roots {
		n.Add(1)
		go walkDir(root, fSizes, &n)
	}

	go func() {
		n.Wait()
		close(fSizes)
	}()

	// Print the results periodically.
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nFiles, nBytes int64

loop:
	for {
		select {
		case size, ok := <-fSizes:
			if !ok {
				break loop // chan was closed
			}
			nFiles++
			nBytes += size
		case <-tick:
			// if -v flag not specified, tick chan is nil, so it's disabled
			fmt.Printf("%d files,  %d bytes\n", nFiles, nBytes)
		}
	}

	fmt.Printf("%d files,  %d bytes\n", nFiles, nBytes)
}
