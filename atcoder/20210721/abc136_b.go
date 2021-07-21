package main

import (
	"fmt"
	"strconv"
)

func stoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func itos(i int) string {
	return strconv.Itoa(i)
}

func main() {
	var n string
	fmt.Scan(&n)
	result := 0
	for i := 0; i < stoi(n); i++ {
		if len(itos(i))%2 == 0 {
			continue
		}
		result++
	}

	if len(n)%2 == 0 {
		result--
	}
	fmt.Println(result)
}
