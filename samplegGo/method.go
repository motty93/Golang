package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func (r *rect) calc() {
	r.width *= 100
	r.height *= 100
}

func (r rect) calc2() {
	r.width *= 100
	r.height *= 100
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())

	// pointerなし
	rp.calc2()
	fmt.Println("calc: ", rp.width)
	fmt.Println("calc: ", rp.height)
	// pointerあり
	rp.calc()
	fmt.Println("calc: ", rp.width)
	fmt.Println("calc: ", rp.height)
}
