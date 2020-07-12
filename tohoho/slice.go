package main

import "fmt"

func display(str string) {
	if str == "red" {
		fmt.Println("Reddddddddddddddddd")
	} else {
		fmt.Println(str)
	}
}

// 可変な配列のことをスライスと呼ぶ
// 不可変な配列よりもメモリ効率や速度は若干落ちる
func main() {
	a1 := []string{}
	// apendで追加
	a1 = append(a1, "red")
	a1 = append(a1, "green")
	a1 = append(a1, "black")
	a1 = append(a1, "blue")

	for i := 0; i < len(a1); i++ {
		display(a1[i])
	}
}
