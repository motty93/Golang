package main

import "fmt"

// interface is any type
func do(i interface{}) {
	// interface型なのでintに変換する
	// ii := i.(int)
	// ii *= 2
	// fmt.Println(ii)

	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T\n", v)
	}
}

func main() {
	do(10)
	do("mike")
	do(false)
}
