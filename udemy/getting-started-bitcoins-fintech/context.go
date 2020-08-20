package main

import (
	"context"
	"fmt"
	"time"
)

// 長くかかりすぎるgoroutineのときにcontextを使う
// contextは渡すだけでおｋ
func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	// ctx.Done()が走る
	// time.Sleep(2 * time.Second)
	time.Sleep(4 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

// ユーザーからのリクエスト処理が長かったりするときによく使われる
func main() {
	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	// とりあえずcontextを設定する場合は以下
	// ctx := context.TODO()
	go longProcess(ctx, ch)
	// cancel()

CTXLOOP:
	for {
		select {
		// 3秒かかって終わらなければこちらが走る
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break CTXLOOP
		case <-ch:
			fmt.Println("success")
			break CTXLOOP
		}
	}
	fmt.Println("####################")
}
