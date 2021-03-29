package main

import "fmt"

func sum(n ...int) int {
	total := 0
	for _, v := range n {
		total += v
	}

	return total
}

func main() {
	a := sum(1, 2, 3)
	fmt.Println(a)

	b := sum(1, 2)
	fmt.Println(b)

	// sliceを渡すときはslice...の形で
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sum(nums...))
}
