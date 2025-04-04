package types

func NewCircle(x, y, r float64) Circle {
	return Circle {
		Center: Point { X: x, Y: y },
		Radius: PrecisionFloat(r),
	}
}

func NewRectangle(x, y, w, h float64) Rectangle {
	return Rectangle {
		TopLeft: Point{X: x, Y: y},
		Width: PrecisionFloat(w),
		Height: PrecisionFloat(h),
	}
}