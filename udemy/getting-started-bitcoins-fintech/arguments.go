package main

import "fmt"

func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}

func main() {
	foo()
	foo(10, 20, 30)

	// 可変長引数としてsliceを渡すこともできる
	s := []int{1, 2, 3}
	fmt.Println(s)
	foo(s...)
}
