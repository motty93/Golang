package main

import (
	"fmt"
	"math/rand"
	"time"
)

const goroutines = 10

const maxProcesses = 3

func main() {
	// セマフォの作成
	semaphore := make(chan int, maxProcesses)
	// 完了通知チャネルの作成
	notify := make(chan int, goroutines)

	for i := 0; i < goroutines; i++ {
		go func(no int, semaphore chan int, notify chan<- int) {
			// 自分の番が来るのを待つ
			semaphore <- 0

			// ダミーの処理として乱数により0-3病態気
			time.Sleep(time.Duration(rand.Int63n(3)) * time.Second)

			fmt.Println("処理完了: ", no)

			// 待機中のゴルーチンに処理を明け渡す
			<-semaphore
			fmt.Println("処理渡すよー")

			// 処理完了を通知
			notify <- no
		}(i, semaphore, notify)
	}

	// ゴルーチンの処理完了待ち
	for j := 0; j < goroutines; j++ {
		num := <-notify
		fmt.Printf("%d番目の処理完了受け取ったよー\n", num)
	}

	// 完了
	fmt.Println("全て完了")
}
