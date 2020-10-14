package main

import "fmt"

func main() {
	name := "taro"
	fmt.Printf("name: %v\n", name)

	// ポインタ型のnamePointを生成
	namePoint := &name
	// ポインタを参照
	fmt.Printf("namePoint :%v\n", namePoint)
	// ポインタの中身を参照
	fmt.Printf("namePoint :%v\n", *namePoint)
	// 中身を変える場合はポインタの中身を指定して代入
	*namePoint = "test"
	fmt.Printf("namePoint :%v\n", *namePoint)
	// ポインタは変わらない
	fmt.Printf("namePoint :%v\n", namePoint)
	// ポインタの中身を書き換えたのでnameも変わる
	fmt.Printf("namePoint :%v\n", name)
}
