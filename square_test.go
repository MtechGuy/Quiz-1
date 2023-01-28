package main

import (
	"testing"
)

func TestSquare(t *testing.T) {
	a, p := square(5)
	if a != 25 {
		t.Errorf("Expected area of 25, got %v", a)
	}
	if p != 20 {
		t.Errorf("Expected perimeter of 20, got %v", p)
	}
}