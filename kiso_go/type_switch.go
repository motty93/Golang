package main

import "fmt"

func showType(x interface{}) {
	switch x.(type) {
	case nil:
		fmt.Println("nil")
	case int, int32, int64:
		fmt.Println("整数")
	case string:
		fmt.Println("文字列")
	default:
		fmt.Println("不明")
	}
}

func showType2(x interface{}) {
	// switchのスコープ内で変数が使用できる
	switch val := x.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("整数", val)
	case string:
		fmt.Println("文字列", val)
	default:
		fmt.Println("不明", val)
	}
}

func main() {
	showType(nil)
	showType("test")
	showType(1000)
	showType(2.13)
	showType2(111)
}
