package main

import "fmt"

func goroutine1(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		ch <- sum
	}
	// 対策 fatal error: all goroutines are asleep - deadlock!
	close(ch)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	// ループを回す長さが決まっているのであれば、channelの長さを指定する
	ch := make(chan int, len(s))
	go goroutine1(s, ch)
	for i := range ch {
		fmt.Println(i)
	}
}
