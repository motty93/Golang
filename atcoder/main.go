package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{}
	// 簡単な標準入力
	fmt.Scan(&a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Printf("入力値：%d\n", a)
}
