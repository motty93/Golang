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
	var n string
	fmt.Scan(&n)
	sum := 0
	s := strings.Split(n, "")
	for i := 0; i < len(s); i++ {
		sum += stoi(s[i])
	}
	if sum%9 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
