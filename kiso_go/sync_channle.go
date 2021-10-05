package main

import (
	"fmt"
	"math/rand"
	"time"
)

const goroutines = 3

func waiting(c chan<- int) {
	// ダミーの処理として乱数により0~10秒待機
	time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

	fmt.Println("処理完了")

	// 処理完了をチャネルで伝える
	// 送信する値は何でも良い
	c <- 10
}

func main() {
	c := make(chan int)

	for i := 0; i < goroutines; i++ {
		go waiting(c)
	}

	// ゴルーチンの処理完了待ち
	for i := 0; i < goroutines; i++ {
		v, ok := <-c

		if !ok {
			break
		}

		fmt.Println(v)
	}

	close(c)
	fmt.Println("全て完了")
}
