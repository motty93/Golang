package main

import (
	"fmt"
	"reflect"
)

// 可変長引数
func holiday(month int, days ...string) {
	fmt.Printf("%d月の祝日は%d日あります。\n", month, len(days))
	for _, day := range days {
		fmt.Println(day)
	}

	fmt.Println()
}

// 可変長をinterface型で待ち受ける
func variable(args ...interface{}) {
	for idx, arg := range args {
		fmt.Println("id：", idx, "型：", reflect.TypeOf(arg), "値：", arg)
	}
}

func main() {
	holiday(1, "元旦", "成人の日")
	holiday(2, "建国記念日")

	variable(1, 2, "test", false)
}
