package main

import (
	"testing"
)

func TestCircleArea(t *testing.T) {
	cir := circle{radius: 5}
	result := cir.area()
	if result != 78.53981633974483 {
		t.Errorf("Expected area of 78.53981633974483, got %v", result)
	}
}

func TestCirclePerimeter(t *testing.T) {
	cir := circle{radius: 5}
	result := cir.perimeter()
	if result != 31.41592653589793 {
		t.Errorf("Expected perimeter of 31.41592653589793")
	}
}
