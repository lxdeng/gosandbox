package main

import (
	"math"
)

type Point2 struct {
	x float64
	y float64
}

// pointer receiver
func (p *Point2) Distance(q Point2) float64 {
	return math.Hypot(q.x-p.x, q.x-p.x)
}
