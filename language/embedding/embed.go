package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	sync.Mutex //lock
	mapping    map[string]int
}

var cache = Cache{mapping: make(map[string]int)}

func main() {
	Put("key1", 100)
	fmt.Printf("%v\n", Lookup("key1"))
}

func Lookup(key string) int {
	cache.Lock() // Lock is promoted from sync.Mutex
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func Put(key string, value int) {
	cache.Lock()
	cache.mapping[key] = value
	cache.Unlock()
}
