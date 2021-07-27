package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	str := "1231231231231asdfasfd"
	for _, s := range str {
		// 型がint32になる
		fmt.Println("型: ", reflect.TypeOf(int(s)))
		fmt.Println(reflect.TypeOf(strconv.Itoa(int(s))))
	}

	var num int32 = 123123123
	fmt.Println(int(num))
}
