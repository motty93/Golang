package main

import "fmt"

type Printable interface {
	ToString() string
}

func PrintOut(p Printable) {
	fmt.Println(p.ToString())
}

type Person struct {
	name string
}

func (p Person) ToString() string {
	return p.name
}

// func (p Person) PrintOut() {
// 	fmt.Println(p.ToString())
// }

type Book struct {
	title string
}

func (b Book) ToString() string {
	return b.title
}

// func (b Book) PrintOut() {
// 	fmt.Println(b.ToString())
// }

func main() {
	a1 := Person{name: "motty"}
	a2 := Book{title: "readable code"}

	// PrintOut関数は同じなので一つにまとめられる
	// a1.PrintOut()
	// a2.PrintOut()
	PrintOut(a1)
	PrintOut(a2)
}
