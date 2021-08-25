package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	var k float64 = 0

	for {
		if int(math.Pow(2.0, k)) > n {
			break
		}

		k++
	}
	fmt.Println(k - 1)
}
