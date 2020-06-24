package main

import "fmt"

type (
	UtcTime string
	JstTime string
)

func main() {
	var (
		t1 UtcTime = "aaaa"
		t2 JstTime = "aaaa"
	)
	// 型変換
	var (
		a1 uint16 = 1234
		a2 uint32 = uint32(a1)
	)

	// type不一致エラー
	fmt.Println("UtcTime value = JstTime value: ", t1 == t2)
}
