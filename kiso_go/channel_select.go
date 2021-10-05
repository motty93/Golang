package main

import (
	"fmt"
	"os"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	// 即時goroutine
	go func() {
		for i := 0; i < 10; i++ {

			// ch1 or ch2へ送信
			select {
			case ch1 <- i:
				fmt.Println("ch1へ送信完了")
			case ch2 <- fmt.Sprintf("test%d\n", i):
				fmt.Println("ch2へ送信完了")
			}
		}

		os.Exit(0)
	}()

	for {
		select {
		case val := <-ch1:
			fmt.Println("ch1からの受信完了：", val)
		case text := <-ch2:
			fmt.Println("ch2からの受信完了：", text)
		}
	}
}
