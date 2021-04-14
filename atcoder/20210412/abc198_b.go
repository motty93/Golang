package main

import "fmt"

func reverseStr(s string) string {
	r := []rune(s)
	n := len(r)
	for i := 0; i < n/2; i++ {
		r[i], r[n-1-i] = r[n-1-i], r[i]
	}
	return string(r)
}

func main() {
	var n string
	fmt.Scan(&n)

	zero := ""
	for j := 0; j < 9; j++ {
		if reverseStr(zero+n) == zero+n {
			fmt.Println("Yes")
			return
		}
		zero += "0"
	}
	fmt.Println("No")
}
