package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int, 10) // initial size
	fmt.Println("m.len=", len(m))
	// cap() is not allowed
	//fmt.Println("m.cap=", cap(m))
	fmt.Printf("%T\n", m)

	m["k1"] = 1
	m["k2"] = 2

	fmt.Println("map: ", m)

	delete(m, "k1")
	fmt.Println("After deletted k1, map: ", m)

	if val, ok := m["k2"]; ok {
		fmt.Printf("key %s exists, value is %d\n\n", "k2", val)
	}

	m["k3"] = 3
	for k, v := range m {
		fmt.Printf("key=%s, value=%d\n", k, v)
	}
}
