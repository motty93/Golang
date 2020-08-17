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

	// 変数に代入しないとキャストできない
	f := 1.11
	fmt.Println(int(f))

	m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 40}
	fmt.Printf("%T %v\n", m, m)
}
