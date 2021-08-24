package main

import "fmt"

func main() {
	array1 := [5]float32{}

	// 不足分はゼロが代入される
	array2 := [6]int{1, 2, 3, 4}

	// 要素数で長さが決まる
	array3 := [...]string{"one", "two", "three"}

	fmt.Println(array1, array2, array3)
}
