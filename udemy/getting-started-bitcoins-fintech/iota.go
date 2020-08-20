package main

import "fmt"

const (
	// 下にもiotaが適応される
	c1 = iota
	c2
	c3
	// c2 = iota
	// c3 = iota
)

const (
	_ = iota
	// 下にもiotaが適応される
	KB int = 1 << (10 * iota)
	MB
	GB
)

func main() {
	fmt.Println(c1, c2, c3)
}
