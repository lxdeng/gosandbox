package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int, 10)
	fmt.Println("m.len=", len(m))
	fmt.Printf("%T\n", m)

	m["k1"] = 1
	m["k2"] = 2

	fmt.Println("map: ", m)

	delete(m, "k1")
	fmt.Println("map: ", m)

	if val, ok := m["k2"]; ok {
		fmt.Printf("key key exists, value is %d\n", val)
	}
}
