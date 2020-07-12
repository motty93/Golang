package main

import "fmt"

func main() {
	// var a1 int
	// var a2 int = 123
	// var a3 = 123
	var (
		a1 int
		a2 int = 123
		a3     = 123
	)

	// 暗黙の型推論初期化
	a4 := 123

	const foo = 100
	const (
		hoge = 123
		huga = 345
	)
}
