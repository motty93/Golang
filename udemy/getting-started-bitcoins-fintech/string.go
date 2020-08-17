package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("hello" + "world")
	// 文字の0番目を取得する
	fmt.Println(string("hello world"[0]))
	var s string = "hello world"

	fmt.Println(strings.Replace(s, "h", "X", 1))
	s = strings.Replace(s, "h", "X", 1)
	fmt.Println(s)
	fmt.Println(strings.Contains(s, "world"))

	fmt.Println(`test
	+ test`)

	fmt.Println("\"")
	fmt.Println(`"`)
}
