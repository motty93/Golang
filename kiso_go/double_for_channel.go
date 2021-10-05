package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5) // バッファ付きchannel宣言

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()

	// 念のためchにデータが書き込まれるのを待つ
	time.Sleep(time.Second)

	// 1〜5を読み込んで出力
	for i := 1; i <= 5; i++ {
		tmp := <-ch
		fmt.Println(tmp)
	}

	// 6〜10がchに書き込まれるのを待つ
	fmt.Println("waiting")
	time.Sleep(time.Second)

	// 上記のループで読み込みが済んで空きが出たので上手く行く
	// 6〜10を読み込んで出力
	for i := 1; i <= 5; i++ {
		tmp := <-ch
		fmt.Println(tmp)
	}
}
