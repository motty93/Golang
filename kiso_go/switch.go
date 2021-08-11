package main

import "fmt"

func main() {
	for i := -2; i <= 2; i++ {
		// trueで固定すればcaseで条件式を使用できる
		// trueは省略可能
		switch true {
		case i > 0:
			fmt.Println("+")
		case i < 0:
			fmt.Println("-")
		default:
			fmt.Println("0")
		}
	}
}
