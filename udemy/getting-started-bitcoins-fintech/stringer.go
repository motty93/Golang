package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("my name is %v.", p.Name)
}

func main() {
	mike := Person{"mike", 20}
	fmt.Println(mike)
}
