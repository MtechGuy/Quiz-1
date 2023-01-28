package main

import (
	"fmt"
	"math"
)

type triangle struct {
	base   float64
	height float64
}
func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func (t triangle) perimeter() float64 {
	a := math.Sqrt(math.Pow(t.base, 2) + math.Pow(t.height, 2))
	return a + t.base + t.height
}