package main

import "fmt"

func addMinus(x int, y int) (int, int) {
	return x + y, x - y
}

func main() {
	// _は参照できない
	// cannnot use _ as value
	_, y := addMinus(10, 12)
	fmt.Println(y)
}
