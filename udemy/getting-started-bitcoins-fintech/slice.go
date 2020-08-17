package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5, 6}
	// endは基本-1
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:4])
	fmt.Println(n[2:])
	fmt.Println(n[:2])
	fmt.Println(n[:])

	n[2] = 1000
	fmt.Println(n[:])

	var board = [][]int{
		[]int{0, 1, 2},
		[]int{0, 1, 2},
		[]int{0, 1, 2},
	}
	fmt.Println(board)

	// 値をappendしたものをnに再代入する
	n = append(n, 100, 200, 300)
	fmt.Println(n)
}
