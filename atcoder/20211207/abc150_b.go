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
	var n int
	var s string
	fmt.Scan(&n, &s)
	fmt.Println(strings.Count(s, "ABC"))
}
