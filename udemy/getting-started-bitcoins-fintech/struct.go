package main

import "fmt"

// Vertex is struct
// 大文字だとそのまま参照できる
// 小文字だとprivate
type Vertex struct {
	X, Y int
	S    string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	// (*v).X = 1000
	// structだと上の書き方を省略して下の書き方でいける
	v.X = 1000
}

func main() {
	v := Vertex{X: 1, Y: 3}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	v.X = 100
	fmt.Println(v.X, v.Y)

	v2 := Vertex{X: 1}
	fmt.Println(v2)
	fmt.Println(v2.S)

	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Println(v4)
	fmt.Printf("%T %v\n", v4, v4)

	// nilにはならない
	var v5 Vertex
	fmt.Println(v5)
	fmt.Printf("%T %v\n", v5, v5)

	// v6/v7は同じ意味
	v6 := new(Vertex)
	fmt.Printf("%T %v\n", v6, v6)
	// 明示的にpointerを取得する場合はこっちにする（よく使われる）
	v7 := &Vertex{}
	fmt.Printf("%T %v\n", v7, v7)

	// map, sliceはmakeが推奨されている

	test := Vertex{1, 2, "test"}
	changeVertex(test)
	fmt.Println(test)

	test2 := &Vertex{1, 2, "test"}
	changeVertex2(test2)
	fmt.Println(*test2)
}
