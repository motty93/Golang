package main

import "fmt"

func main() {
	n := 0
	for {
		n++
		if n > 10 {
			break
		} else if n%2 == 1 {
			continue
		} else {
			fmt.Println(n)
		}
	}

	// 配列・スライス・mapなどのイテラブルな物に対するforループ
	colors := [...]string{"red", "green", "blue"}
	for i, color := range colors {
		fmt.Printf("%d: %s\n", i, color)
	}

	foods := map[int]string{
		10: "apple",
		11: "orange",
	}
	for k, v := range foods {
		fmt.Printf("%d=%s\n", k, v)
	}
}
