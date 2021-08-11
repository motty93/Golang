package main

import "fmt"

func hoge(i int) (int, int) {
	return i + 1, i - 1
}

func main() {
	if x, y := hoge(2); x < y {
		fmt.Println(x, "<", y)
	} else if x > y {
		fmt.Println(x, ">", y)
	} else {
		fmt.Println(x, "=", y)
	}
}
