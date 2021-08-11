package main

import "fmt"

func main() {
	// intポインタ変数
	var ptr *int

	var i int = 12345

	// アドレス演算子を使いiのアドレスをptrに代入
	ptr = &i

	fmt.Println("iのあどれす：", &i)
	fmt.Println("ptrの値：", ptr)

	fmt.Println("iの値：", i)
	fmt.Println("ポインタ経由のiの値：", *ptr)

	// ポインタ経由でiの値を変更
	*ptr = 9876
	fmt.Println("iの値：", i)
	fmt.Println("ポインタ経由のiの値：", *ptr)
}
