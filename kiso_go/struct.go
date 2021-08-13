package main

import "fmt"

// 構造体に名前をつける（一般的にこっちを使う）
type Data struct {
	s string
	b byte
	_ byte // ブランクフィールド
}

// 匿名フィールド
// それぞれのフィールド名は型の名前になる
// 埋め込みと呼ばれる
struct {
	int
	string
	bool
}

func main() {
	var x struct {
		i int
		f float32
		s string
		r rune
	}

	x.i = 1
	x.f = 0.1
	x.s = "test"

	fmt.Println(x)
}
