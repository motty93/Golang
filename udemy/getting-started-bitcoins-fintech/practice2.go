package main

import (
	"fmt"
	"sort"
)

func minNumber(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.SliceStable(nums,
		func(i, j int) bool {
			return nums[i] < nums[j]
		})

	return nums[0]
}

func main() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6}
	m := []int{}
	fmt.Println(minNumber(l...))
	fmt.Println(minNumber(m...))
}
