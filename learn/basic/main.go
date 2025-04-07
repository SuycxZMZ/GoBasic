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

	// 接口类型的切片，可以存储任何实现了 Shape 接口的类型的实例。
	// 这使得我们可以在不关心具体类型的情况下，对这些类型进行操作。
	// 例如，我们可以遍历这个切片，调用每个类型的 Area 和 Perimeter 方法。
	// 这使得我们可以编写通用的代码，来处理不同类型的形状。
	shape := []types.Shape{circle, rect}

	for _, s := range shape {
		printShapeInfo(s)

		// if initialization; condition { ... }
		// 类型断言
		// 类型断言用于将接口类型转换为具体类型。
		// 例如，我们可以使用类型断言将接口类型转换为 Circle 类型，然后调用 Circle 类型的方法。
		// 类型断言的语法是：
		// value, ok := interfaceValue.(ConcreteType)
		if c, ok := s.(types.Circle); ok {
			fmt.Printf("Circle radius: %f\n", c.Radius)
		}
	}

	var f types.PrecisionFloat = 3.1415
	fmt.Printf("PrecisionFloat: %.4f\n", float64(f))
}