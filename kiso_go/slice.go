package main

import "fmt"

func double(values []int) {
	for i := 0; i < len(values); i++ {
		values[i] *= 2
	}
}

func main() {
	// sliceは参照型の一種
	s1 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(s1[:4])

	// pointerを渡さずとも値を変更できる
	s2 := []int{0, 2, 3, 4, 5}
	double(s2)
	fmt.Println(s2)
}
