package main

import "fmt"

func main() {
	// int型からmyInteger方を宣言
	type myInteger int

	var i myInteger = 123
	i++
	fmt.Println(i)

	type myStruct struct {
		a, b int
	}
}
