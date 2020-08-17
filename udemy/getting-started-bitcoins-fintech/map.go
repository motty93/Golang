package main

import "fmt"

func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 100000
	fmt.Println(m)
	m["new"] = 200
	fmt.Println(m)

	// 0が返ってくる
	fmt.Println(m["nothing"])

	v, ok := m["apple"]
	if ok != false {
		fmt.Println(v)
	}

	v2, ok2 := m["nothing"]
	if ok2 != false {
		fmt.Println(v2)
	}

	m2 := make(map[string]int)
	m2["pc"] = 20000
	fmt.Println(m2)

	// メモリ上にないものに値を入れようとしてるのでpanic
	// var m3 map[string]int
	// m3["pc"] = 2000
	// fmt.Println(m3)

	// varで宣言するとmapでもsliceでもnilになる
	var s []int
	if s == nil {
		fmt.Println("nil")
	}
}
