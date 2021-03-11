package main

import "fmt"

var x int

func main() {
	fmt.Scan(&x)
	if x%100 == 0 {
		fmt.Println(100)
	} else {
		fmt.Printf("%d\n", 100-x%100)
	}
}
