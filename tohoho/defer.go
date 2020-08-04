package main

import "os"

func main() {
	fp, err := os.Open("sample.txt")
	if err != nil {
		return
	}

	// リソースの解放を遅延実行
	defer fp.Close()
}
