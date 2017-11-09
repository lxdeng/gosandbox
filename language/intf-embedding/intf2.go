package main

import (
	"fmt"
)

type Walker interface {
	Walk(n int)
}

type Runner interface {
	Run(n int)
}

type Mover interface {
	Walker
	Runner
}

type Dog struct {
	name string
}

func main() {
	var m Mover = &Dog{"tom"}
	m.Walk(1)
	m.Run(2)
}

func (d *Dog) Walk(n int) {
	fmt.Printf("%s walked %d miles\n", d.name, n)
}

func (d *Dog) Run(n int) {
	fmt.Printf("%s run %d miles\n", d.name, n)
}
