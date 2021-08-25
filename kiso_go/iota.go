package main

import "fmt"

const (
	// iotaは列挙子で呼び出されるたびauto incrementされる
	ZERO  = iota
	ONE   = iota
	TWO   = iota
	THREE = iota
)

func main() {
	fmt.Println(ZERO)
	fmt.Println(ONE)
	fmt.Println(TWO)
	fmt.Println(THREE)
}
