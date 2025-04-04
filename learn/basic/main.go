package main

import (
	"fmt"
	"learn/basic/types"
)

func printShapeInfo(s types.Shape) {
	fmt.Printf("Area: %f\n Perimeter:%.2f", s.Area(), s.Perimeter())
}

func main() {
	circle := types.NewCircle(0, 0, 5)
	rect := types.NewRectangle(0, 0, 4, 6)

	shape := []types.Shape{circle, rect}

	for _, s := range shape {
		printShapeInfo(s)
		if c, ok := s.(types.Circle); ok {
			fmt.Printf("Circle radius: %f\n", c.Radius)
		}
	}

	var f types.PrecisionFloat = 3.1415
	fmt.Printf("PrecisionFloat: %.4f\n", float64(f))
}