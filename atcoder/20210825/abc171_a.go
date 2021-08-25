package main

import (
	"fmt"
	"unicode"
)

func main() {
	var a string
	fmt.Scan(&a)
	r := rune(a[0])

	if unicode.IsUpper(r) {
		fmt.Println("A")
	} else {
		fmt.Println("a")
	}
}
