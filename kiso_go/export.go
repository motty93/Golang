package main

// 先頭が大文字なのでexportされる
const Export = true

// 小文字のためexportされない
const export = false

func main() {
	// ローカル変数のためexportされない
	const A = 123
}
