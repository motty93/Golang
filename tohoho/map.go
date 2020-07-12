package main

import "fmt"

func main() {
	a1 := map[string]int{
		"x": 0,
		"b": 100,
	}
	// 参照
	fmt.Println(a1["x"])

	// マップに要素を加える
	a1["z"] = 3000

	// マップの要素を削除
	delete(a1, "b")

	// 長さ
	fmt.Println(len(a1))

	// マップに要素が存在するかどうかを調べる
	// マップのキーに参照すると、値と存在有無のbool値が戻り値になる
	_, ok := a1["z"]
	if ok {
		fmt.Println("exist")
	} else {
		fmt.Println("not exist")
	}
	// 短く書くなら
	if _, ok := a1["xxxx"]; ok {
		fmt.Println("exist")
	} else {
		fmt.Println("not exist.")
	}

	// マップのループ
	for k, v := range a1 {
		fmt.Printf("%s=%d\n", k, v)
	}
}
