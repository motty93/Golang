package main

import "fmt"

// Printable is common type interface
type Printable interface {
	ToString() string
}

// PrintOut is common function insterface with any
func PrintOut(a interface{}) {
	// aをPrintableインタフェースを実装したオブジェクトに変換してみる
	q, ok := a.(Printable)
	if ok {
		// 変換できてたらそのインタフェースを呼び出す
		fmt.Println(q.ToString())
	} else {
		fmt.Println("Not printable")
	}
}

// switch 変数.(type) {...}で型を判断することもできる
func test(a interface{}) {
	switch a.(type) {
	case bool:
		fmt.Printf("%t\n", a)
	case int:
		fmt.Printf("%d\n", a)
	case string:
		fmt.Printf("%s\n", a)
	}
}

// Person is struct
type Person struct {
	name string
}

// ToString is Person function
func (p Person) ToString() string {
	return p.name
}

// Article is struct
type Article struct {
	name string
}

func main() {
	a1 := Person{name: "motty"}
	a2 := Article{name: "Biology"}
	PrintOut(a1)
	PrintOut(a2)

	p1 := map[string]interface{}{
		"name": "motty",
		"age":  90,
	}
}
