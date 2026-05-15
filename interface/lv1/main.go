package main

type Animal interface {
	Sound() string
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func (d Dog) Sound() string {
	return "Baw!"
}

func (c Cat) Sound() string {
	return "Meow!"
}

func main() {
	cat := Cat{name: "kitty"}
	dog := Dog{name: "john"}

	animals := []Animal{cat, dog}
	for _, a := range animals {
		println(a.Sound())
	}
}
