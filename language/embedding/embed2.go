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
	cache.Put("key1", 100)
	fmt.Printf("%v\n", cache.Lookup("key1"))
}

func (c Cache) Lookup(key string) int {
	c.Lock() // Lock is promoted from sync.Mutex
	v := cache.mapping[key]
	c.Unlock()
	return v
}

func (c Cache) Put(key string, value int) {
	c.Lock()
	c.mapping[key] = value
	c.Unlock()
}
