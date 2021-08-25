package main

import "fmt"

// int型から当たらにmyType型を宣言
type myType int

func (value myType) println() {
	fmt.Println(value)
}

func (value myType) setByValue(newValue myType) {
	value = newValue
}

func (value *myType) setByPonter(newValue myType) {
	*value = newValue
}

func main() {
	var z myType = 10
	z.println()

	var x myType = 0
	// レシーバがタダの値なので変更不可
	x.setByValue(1)
	fmt.Println(x)

	// レシーバがポインタなので変更可
	x.setByPonter(2)
	fmt.Println(x)
}
