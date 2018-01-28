package main

import (
	"fmt"
)

func main() {
	p := Point{1, 1}
	q := Point{2, 2}

	// same type, Point
	fmt.Println(p.Distance(q))

	pp := &p
	// compiler implicitly dereferences the receiver, i.e. loads the value
	fmt.Println(pp.Distance(q))

	p2 := Point2{1, 1}
	q2 := Point2{2, 2}

	// same type, *Point2
	fmt.Println((&p2).Distance(q2))

	// compiler implicitly uses the address of the variable p2
	fmt.Println(p2.Distance(q2))

	//the following does not compile, since compiler cannot take the address of the literal
	//Point{1, 2}.Distance(q2)
}
