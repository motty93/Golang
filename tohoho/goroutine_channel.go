package main

import (
	"fmt"
	"time"
)

func funcA(chA chan<- string) {
	time.Sleep(3 * time.Second)
	chA <- "finished"
}

func main() {
	chA := make(chan string) // チャネルの作成
	defer close(chA)         // 使い終わったらクローズ
	go funcA(chA)            // チャネルをゴルーチンに渡す
	msg := <-chA             // チャネルからメッセージを受信する

	fmt.Println(msg)
}
