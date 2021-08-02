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
	var xn string
	fmt.Scan(&xn)
	xi := strings.Split(xn, "")
	uniq := make(map[string]bool)
	for _, x := range xi {
		if !uniq[x] {
			uniq[x] = true
		}
	}
	if len(uniq) == 1 {
		fmt.Println("Weak")
		return
	}

	for idx, x := range xi {
		if len(xi) == idx+1 {
			break
		}

		if stoi(x) == 9 {
			if stoi(xi[idx+1]) != 0 {
				fmt.Println("Strong")
				return
			}
		} else {
			if stoi(x)+1 != stoi(xi[idx+1]) {
				fmt.Println("Strong")
				return
			}
		}
	}

	fmt.Println("Weak")
}
