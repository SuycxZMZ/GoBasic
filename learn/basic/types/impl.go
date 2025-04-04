package types

import "math"

type Circle struct {
	Center Point
	Radius PrecisionFloat
}

func (c Circle) Area() float64 {
	return math.Pi * float64(c.Radius) * float64(c.Radius)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * float64(c.Radius)
}

type Rectangle struct {
	TopLeft  Point
	Width, Height PrecisionFloat
}

func (r Rectangle) Area() float64 {
	return float64(r.Width) * float64(r.Height)
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (float64(r.Width) + float64(r.Height))
}

