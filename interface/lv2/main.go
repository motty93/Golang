package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectagle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectagle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func PrintArea(s Shape) {
	fmt.Printf("area: %.2f\n", s.Area())
}

func main() {
	r := Rectagle{Width: 10, Height: 20}
	c := Circle{Radius: 5}
	shapes := []Shape{r, c}

	for _, s := range shapes {
		PrintArea(s)
	}
}
