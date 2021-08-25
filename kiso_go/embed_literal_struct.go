package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id uint
	Person
}

func main() {
	e := Employee{1, Person{"motty", 30}}

	fmt.Println(e)
}
