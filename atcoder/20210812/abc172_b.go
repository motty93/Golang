package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	count := 0
	for idx, v := range strings.Split(s, "") {
		fmt.Println(t[idx : idx+1])
		if t[idx:idx+1] != v {
			count++
		}
	}
	fmt.Println(count)
}
