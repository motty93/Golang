package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(s[1:] + s[:1])
}
