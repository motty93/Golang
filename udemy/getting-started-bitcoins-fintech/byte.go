package main

import "fmt"

func main() {
	b := []byte{90, 9}
	fmt.Println(b)
	fmt.Println(string(b))

	c := []byte("hi")
	fmt.Println(c)
	fmt.Println(string(c))
}
