package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("status:", res.Status)

	// res bodyを全て読み込む
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(body))
}
