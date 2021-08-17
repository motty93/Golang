package main

import (
	"fmt"
)

func main() {
	var s, t int
	fmt.Scan(&s, &t)

	count := 0
	for a := 0; a < s+1; a++ {
		for b := 0; b < s+1-a; b++ {
			for c := 0; c < s+1-c-b; c++ {
				if a+b+c <= s && a*b*c <= t {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
