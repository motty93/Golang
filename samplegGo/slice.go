package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp: ", s)
	fmt.Println("len: ", len(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)

	// 破壊的に追加される
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	c := make([]string, len(s))
	copy(c, s) // コピー可能
	fmt.Println("cp: ", c)

	// index2から5の前まで取得
	l := s[2:5]
	fmt.Println("sl1: ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t)

	// 二次元slice
	twoD := make([][]int, 3)
	fmt.Println(twoD)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
