package main

import (
	"fmt"
	"os"
)

const goroutines = 10

func main() {
	counter := make(chan int, 1)
	// counter channelに初期値を送信
	counter <- 0

	for i := 0; i < goroutines; i++ {
		go func(counter chan int) {
			// チャネルから値を取り出す
			value := <-counter

			value++

			fmt.Println("counter:", value)

			// 全てのゴルーチンの処理が完了したら終了
			if value == goroutines {
				os.Exit(0)
			}

			counter <- value
		}(counter)
	}

	// 無限ループ
	for {
	}
}
