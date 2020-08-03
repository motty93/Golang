package main

import "fmt"

type Person struct {
	Name   string // 外部アクセス可
	Age    int    // 外部アクセス可
	name   string // 外部アクセス不可
	age    int    // 外部アクセス不可
	status int    // 外部アクセス不可
}

func (p *Person) SetPerson(name string, age int) {
	p.Name = name
	p.name = name
	p.age = age
}

func (p *Person) GetPerson() (string, int) {
	return p.name, p.age
}

func main() {
	var (
		person Person
		name   string
		age    int
	)

	person.SetPerson("motty", 29)
	name, age = person.GetPerson()
	fmt.Println(name, age)
	fmt.Println(person.Name) // 外部アクセス可

	// 順序通りに初期化
	a1 := Person{"motty", 93, "motty", 20, 1}
	// 名前で初期化
	a2 := Person{name: "tanaka", age: 20, Name: "tanaka", Age: 20, status: 10}

	fmt.Println(a1.Name, a1.Age)
	fmt.Println(a2.Name, a2.Age)
}
