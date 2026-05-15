package main

import (
	"fmt"
	"reflect"
)

type ErrorCode int

const (
	hoge int       = 100
	code ErrorCode = 10
)

func main() {
	fmt.Println("type check start")

	a := hoge
	fmt.Println(reflect.TypeOf(a))

	b := code
	fmt.Println(reflect.TypeOf(b))

	// type error
	// a = b
}
