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
