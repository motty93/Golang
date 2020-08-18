package main

import "fmt"

type Vertex struct {
	x, y int
}

// 値レシーバ
func (v Vertex) Area() int {
	return v.x * v.y
}

// func Area(v Vertex) int {
// 	return v.x * v.y
// }

// pointerレシーバ
// pointerにすると中身を書き換えることができる
func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

// golangのデザインパターン
func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

func main() {
	// fmt.Println(Area(v))
	// v := Vertex{3, 4}
	v := New(3, 4)
	v.Scale(10)
	fmt.Println(v.Area())
}
