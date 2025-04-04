package types

import (
	"math"
	"testing"
)

func TestCircleArea(t * testing.T) {
	c := Circle{ Radius: 5 }
	expect := math.Pi * 5 * 5
	if math.Abs(c.Area() - expect) > 1e-9 {
		t.Errorf("Expected area %.4f, got %.4f", expect, c.Area())
	}
}

func TestRectanglePerimeter(t * testing.T) {
	r := Rectangle{Width: 4, Height: 6}
	expected := 2 * (4 + 6)
	if r.Perimeter() != float64(expected) {
		t.Errorf("Expected perimeter %.4f, got %.4f", float64(expected), r.Perimeter())
	}
}