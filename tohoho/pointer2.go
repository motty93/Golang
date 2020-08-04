package main

import "fmt"

// ChangePointer is function
func ChangePointer(a1 int, a2 *int) {
	a1 = 1000000
	*a2 = 1000000
}

// Person is struct
type Person struct {
	name string
	age  int
}

func main() {
	a1 := 123
	a2 := 456
	fmt.Println(&a2)
	fmt.Println("===========================")
	ChangePointer(a1, &a2) // a1は値渡し、a2は参照渡し
	// &a2にするとrubyのインスタンス変数のように使える
	fmt.Println(&a2)
	fmt.Println(a2)

	p1 := Person{"motty", 70}
	q1 := &p1
	fmt.Println(p1.name)

	// 下2つは同じ意味
	fmt.Println((*q1).name)
	fmt.Println(q1.name)
}
