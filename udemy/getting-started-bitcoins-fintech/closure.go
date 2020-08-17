package main

import "fmt"

// 関数をreturnする
func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func sum(x, y int) int {
	return x + y
}

func partialSum() func() int {
	return func() (x int) {
		x = 100 + 200
		return
	}
}

func main() {
	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	c1 := circleArea(3.14)
	fmt.Println(c1(2))
	fmt.Println(c1(3))

	partial := partialSum()
	fmt.Println(partial())
}
