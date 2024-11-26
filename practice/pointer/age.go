package main

import "fmt"

type AgeUpdater interface {
	UpdateAge(age int)
	GetAge() int
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) UpdateAge(age int) {
	p.Age = age
}

func (p *Person) GetAge() int {
	return p.Age
}

type Animal struct {
	Species string
	Age     int
}

func (a *Animal) UpdateAge(age int) {
	a.Age = age
}

func (a *Animal) GetAge() int {
	return a.Age
}

func main() {
	p := &Person{"motty", 33}
	a := &Animal{"dog", 5}
	// 通常のudpate age
	p.UpdateAge(40)
	fmt.Println("p.UpdateAge:", p.Age) // 40になる

	updaters := []AgeUpdater{p, a}
	// 全てのインターフェース型を操作
	for _, u := range updaters {
		currentAge := u.GetAge()
		currentAge++
		u.UpdateAge(currentAge)
	}

	for _, u := range updaters {
		fmt.Println(u)
	}
}
