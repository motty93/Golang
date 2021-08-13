package main

import (
	"fmt"
	"os"
)

func delay() {
	fmt.Println("delayが呼び出されました")
}

func fileOpen() {
	// ファイルの作成書き込み
	file, err := os.OpenFile("text.txt", os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		os.Exit(1)
	}

	defer file.Close()

	file.WriteString("あいうえお")
}

func main() {
	fmt.Println("開始")

	// 遅延実行した関数からは戻り地を受け取れない
	// 関数内の処理が終わってから実行される
	defer delay()
	fileOpen()

	fmt.Println("終了")
}
