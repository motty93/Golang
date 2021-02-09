package main

import (
	"fmt"
	"strconv"
)

var s string

func main() {
	fmt.Scan(&s)
	i, _ := strconv.Atoi(s[:2])
	j, _ := strconv.Atoi(s[2:])

	switch {
	case i == 0 || j == 0:
		fmt.Println("NA")
	case j >= 1 && j <= 12:
		fmt.Println("YYMM")
	case i >= 1 && i <= 12:
		fmt.Println("MMYY")
	default:
		fmt.Println("AMBIGUOUS")
	}
}
