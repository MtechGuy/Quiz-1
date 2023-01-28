package main

import (
	"testing"
)

func TestTriangleArea(t *testing.T) {
	tri := triangle{base: 3, height: 4}
	result := tri.area()
	if result != 6 {
		t.Errorf("Expected area of 6, got %v", result)
	}
}

func TestTrianglePerimeter(t *testing.T) {
	tri := triangle{base: 3, height: 4}
	result := tri.perimeter()
	if result != 12 {
		t.Errorf("Expected perimeter of 12, got %v", result)
	}
}
