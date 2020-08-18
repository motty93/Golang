package main

import (
	"fmt"
	"os"
)

func foo() {
	// function fooの最後に実行される
	defer fmt.Println("world defer")

	fmt.Println("hello defer")
}

func main() {
	// deferはmain処理の最後に実行される
	// defer fmt.Println("world foo")
	// defer foo()
	//
	// fmt.Println("hello fooo")

	// 最初にdeferしたものが最後にくる
	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")

	// ファイル読み込み
	file, _ := os.Open("./byte.go")
	// 処理が終わったら閉じる
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}
