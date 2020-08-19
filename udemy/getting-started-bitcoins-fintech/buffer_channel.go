package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))
	close(ch)

	// 3つめのchを撮ろうとするのでerrorになる
	// rangeでforループする場合、その前にcloseが必要
	for c := range ch {
		fmt.Println(c)
	}
}
