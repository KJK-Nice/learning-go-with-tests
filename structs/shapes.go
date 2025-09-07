package structs

import "math"

type Shape interface {
    Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}


func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circles struct {
    Radius float64
}

func (c Circles) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
