package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var s string = "14"
	// 文字列の数字変換
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v %d\n", i, i, i)

	h := "Hello world"
	fmt.Println(string(h[0]))
}
