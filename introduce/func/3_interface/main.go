package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func PrintArea(s Shape) {
	fmt.Printf("Area = %.2f\n", s.Area())
}

func main() {
	r := Rectangle{Width: 5, Height: 3}
	c := Circle{Radius: 4}
	
	PrintArea(r)
	PrintArea(c)
}
