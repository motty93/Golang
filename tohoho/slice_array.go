package main

import "fmt"

func main() {
	a := [5]int{2, 3, 4, 5, 6}

	b := a[2:4]
	c := a[2:]
	d := a[:4]
	e := a[:]
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// make関数でいきなりスライスを生成する
	s1 := make([]int, 3)

	// いきなり値を割り当てたスライスを生成する
	s2 := []int{1, 2, 3}
	fmt.Println(s1)
	fmt.Println(s2)

	// appendでスライスの末尾に要素を追加
	s3 := append(s2, 6, 4, 5)
	fmt.Println(s3)
	// appendは破壊的ではない
	fmt.Println(s2)

	// copyでスライスをコピー
	// 返り値は要素数
	t := make([]int, len(s3))
	s4 := copy(t, s3)
	fmt.Println(s4)
	fmt.Println(t)
}
