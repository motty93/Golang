package main

import (
	"fmt"
	"strconv"
	"strings"
)

func stoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func main() {
	var n, s string
	fmt.Scan(&n, &s)
	if strings.Index(s, "1")%2 == 0 {
		fmt.Println("Takahashi")
	} else {
		fmt.Println("Aoki")
	}
}
