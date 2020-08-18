package main

import "fmt"

func one(x *int) {
	// xのポインターを受け取って実体に対して20を代入する
	*x = 20
}

func main() {
	var n int = 100
	one(&n)
	fmt.Println(n)
	// fmt.Println(n)
	// fmt.Println(&n)
	//
	// var p *int = &n
	// fmt.Println(p)
	// fmt.Println(*p)
}
