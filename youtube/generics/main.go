package main

import "fmt"

type Number interface {
	int8 | int16 | int32 | int64 | float32 | float64
}

// generics
func newGenericFunc[age Number](myAge age) {
	value := float32(myAge) + 1
	fmt.Println(value)
}

// pointerを渡す
func BubbleSort[N Number](input *[]N) {
	n := len(*input)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if (*input)[i] > (*input)[i+1] {
				(*input)[i], (*input)[i+1] = (*input)[i+1], (*input)[i]
				swapped = true
			}
		}
	}
}

func main() {
	fmt.Println("Hello, World!")

	var testAge int64 = 28
	var testAge2 float32 = 28.6

	newGenericFunc(testAge)
	newGenericFunc(testAge2)

	list := []int32{5, 3, 2, 1, 4}
	BubbleSort(&list)
	fmt.Println(list)

	listFloat := []float32{5.5, 3.3, 2.2, 1.1, 4.4}
	BubbleSort(&listFloat)
	fmt.Println(listFloat)
}
