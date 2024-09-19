package shapes

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

func (r Circle) Area() float64 {
	return math.Pi * r.Radius * r.Radius
}
