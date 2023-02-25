package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{5, 6}

	s2 = append(s2, 7, 8)
	fmt.Println(s2)

	// sliceにsliceをappend
	s1 = append(s1, s2...)
	fmt.Println(s1)
}