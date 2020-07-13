package main

import "fmt"

func main() {
	n := 0
	for {
		n++
		if n > 10 {
			break
		} else if n%2 == 1 {
			continue
		} else {
			fmt.Println(n)
		}
	}
}
