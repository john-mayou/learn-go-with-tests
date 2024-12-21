package shapes

import "math"

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type RightTriangle struct {
	Base   float64
	Height float64
}

func (r *RightTriangle) Perimeter() float64 {
	a := r.Height
	b := r.Base
	c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	return a + b + c
}

func (r *RightTriangle) Area() float64 {
	return (r.Base * r.Height) / 2
}
