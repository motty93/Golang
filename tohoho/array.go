package main

import "fmt"

// 個数が決まっていて変更不可のものを配列とする
func main() {
	a1 := [3]string{}
	a1[0] = "red"
	a2[1] = "blue"
	a3[2] = "green"

	a10 := [3]string{"red", "blue", "green"}
	// 初期化で個数が決まる場合は...と省略可能(3と明示してもおｋ)
	b1 := [...]string{"red", "blue", "green"}

	fmt.Println(a1[0], a1[1], a1[2])
}
