package main

import "fmt"

// objectのポインタ参照&
// ポインタの中身を参照*
func main() {
	var (
		a1 int  // int型変数a1を定義
		p1 *int // intへのポインタ変数p1を定義
	)

	// p1にa1のポインタを代入
	a1 = 100
	p1 = &a1
	fmt.Println(a1)
	fmt.Println(p1)
	fmt.Println(*p1)
	fmt.Println("===============================")
	*p1 = 200
	fmt.Println(a1)
}
