package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 12
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "alice", age: 12})

	fmt.Println(&person{name: "ann"})

	fmt.Println(*newPerson("John"))
	fmt.Println(newPerson("Point"))

	p := person{name: "moto", age: 20}
	fmt.Println(p.name)

	sp := &p
	fmt.Println(sp.age)
	sp.age = 100
	fmt.Println(sp.age)
	fmt.Println(p.age)
}
