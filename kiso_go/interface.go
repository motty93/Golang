package main

import "fmt"

type Calculator interface {
	Calculate(a, b int) int
}

type Add struct{}

func (x Add) Calculate(a, b int) int {
	return a + b
}

type Sub struct{}

func (x Sub) Calculate(a, b int) int {
	return a - b
}

func main() {
	m := map[interface{}]interface{}{
		1:      10,
		"test": 20,
	}

	fmt.Println(m)

	// Calculatorインターフェースを実装した型の変数を宣言
	var add Add
	var sub Sub

	// Calculatorインターフェース型宣言
	var cal Calculator
	// Add型を代入
	cal = add
	fmt.Println("和：", cal.Calculate(1, 2))

	// Sub型を代入
	cal = sub
	fmt.Println("差：", cal.Calculate(1, 2))
}
