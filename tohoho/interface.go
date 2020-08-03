package main

import "fmt"

// Printable common PrintOut function interface
type Printable interface {
	ToString() string
}

// PrintOut is Println with function of ToString
func PrintOut(p Printable) {
	fmt.Println(p.ToString())
}

// Person is struct
type Person struct {
	name string
}

// ToString return Person name
func (p Person) ToString() string {
	return p.name
}

// func (p Person) PrintOut() {
// 	fmt.Println(p.ToString())
// }

// Book is struct
type Book struct {
	title string
}

// ToString is Book struct class method
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
