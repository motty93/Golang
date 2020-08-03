package main

import "fmt"

func addMinus(x int, y int) (int, int) {
	return x + y, x - y
}

// 可変引数を渡せる
func funcA(a int, b ...int) {
	fmt.Printf("a=%d\n", a)
	for i, num := range b {
		fmt.Printf("b[%d] = %d\n", i, num)
	}
}

func main() {
	// _は参照できない
	// cannnot use _ as value
	_, y := addMinus(10, 12)
	fmt.Println(y)
	funcA(1, 2, 3, 4, 5)
}
