package types

// import "internal/itoa"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type PrecisionFloat float64

type Color int 

// const (
// 	Red Color = itoa
// 	Green
// 	Blue
// )

type Point struct {
	X, Y float64
}