package main

import "fmt"

func multiply(x, y int) int {
	return x * y
}

func main() {
	// 関数型の宣言
	var f func(int, int) int

	// 関数型の変数に関数リテラルの値を代入
	// jsのarrow functionみたいな感じにできる
	f = func(a, b int) int {
		return a + b
	}

	// 関数型の変数経由で関数を呼び出す
	fmt.Println(f(1, 2))

	f = multiply
	fmt.Println(f(1, 2))

	// こんな感じでつかう
	hoge := func(a int) int {
		return a * 10
	}
	fmt.Println(hoge(10))
}
