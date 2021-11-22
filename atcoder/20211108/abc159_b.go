package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	fmt.Scan(&s)
	b := s[:(len(s)-1)/2]
	a := s[(len(s)-1)/2+1:]
	if b != a {
		fmt.Println("No")
		os.Exit(0)
	}

	for idx, v := range b {
		if !(string([]rune{v}) == b[len(b)-(idx+1):len(b)-idx]) {
			fmt.Println("No")
			os.Exit(0)
		}

		if (len(b)-1)/2 == idx {
			break
		}
	}

	fmt.Println("Yes")
}
