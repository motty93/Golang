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

	// capは容量を求める
	fmt.Println(cap(a1))

	// makeを用いてメモリの確保が可能。
	// 容量超過時の再確保を減らして速度を早めることができる
	fmt.Println("========a2========")
	a2 := make([]string, 0, 1024)
	a2 = append(a2, "red")
	a2 = append(a2, "green")
	a2 = append(a2, "black")
	a2 = append(a2, "blue")

	for i := 0; i < len(a2); i++ {
		display(a2[i])
	}
}
