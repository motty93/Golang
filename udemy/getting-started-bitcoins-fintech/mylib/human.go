package mylib

import "fmt"

var Public string = "public"
var private string = "private"

type Person struct {
	Name string
	Age  int
}

// Stringer defined
func (p Person) String() string {
	return fmt.Sprintf("my name is %s", p.Name)
}

func Say() {
	fmt.Println("hello")
}
