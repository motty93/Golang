package main

import (
	"fmt"
	"time"
)

func test() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		// 1秒待つ
		time.Sleep(1 * time.Second)
	}
}

func main() {
	fmt.Println("main:開始")
	fmt.Println("testを通常関数として起動")
	test()

	fmt.Println("testをゴルーチンとして起動")
	go test()

	// 3秒待つ
	time.Sleep(3 * time.Second)

	fmt.Println("main:終了")
}
