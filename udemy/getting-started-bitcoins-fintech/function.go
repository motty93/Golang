package main

import "fmt"

// resultが返ることを最初に宣言しているのでreturn指定する必要なし
func cal(price, item int) (result int) {
	result = price * item
	// return result
	return
}

func main() {
	r := cal(100, 30)
	fmt.Println(r)

	f := func(x int) {
		fmt.Println("hello inner func", x)
	}
	f(1)

	// 即時関数的なやつ
	func(x int) {
		fmt.Println("hello inner func", x)
	}(2000)
}
