package main

import (
	"math"
)

type Point struct {
	x float64
	y float64
}

// value (non-pointer) receiver
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.x-p.x)
}
