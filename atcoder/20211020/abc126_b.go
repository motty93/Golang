package main

import (
	"fmt"
	"os"
	"strconv"
)

func stoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func main() {
	var s string
	fmt.Scan(&s)
	bes := stoi(s[:2])
	afs := stoi(s[2:4])
	if bes == 0 || afs == 0 || (bes > 12 && afs > 12) {
		fmt.Println("NA")
		os.Exit(0)
	}

	switch {
	case bes <= 12 && afs <= 12:
		fmt.Println("AMBIGUOUS")
	case bes <= 12 && afs > 12:
		fmt.Println("MMYY")
	case bes > 12 && afs <= 12:
		fmt.Println("YYMM")
	default:
		fmt.Println("NA")
	}
}
