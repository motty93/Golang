package main

import (
	"fmt"
	"time"
)

func main() {
	// 5秒ごとに通知されるチャネルを作成
	ch := time.Tick(time.Second * 5)

	for i := 0; i < 10; i++ {
		t := <-ch
		fmt.Println(t)
	}
}
