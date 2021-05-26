package main

import (
	"fmt"
	"strings"
)

func uniqueSlice(target []string) (unique []string) {
	m := map[string]bool{}

	for _, v := range target {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}

func main() {
	var s string
	fmt.Scan(&s)
	t := strings.Split(s, "")
	r := uniqueSlice(t)
	if len(r) == 2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
