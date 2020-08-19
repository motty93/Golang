package mylib

import "fmt"

var Public string = "public"
var private string = "private"

type Person struct {
	Name string
	Age  int
}

func Say() {
	fmt.Println("hello")
}
