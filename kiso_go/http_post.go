package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	values := url.Values{}
	values.Add("title", "foo")
	values.Add("body", "bar")
	values.Add("userId", "1")
	res, err := http.PostForm("https://jsonplaceholder.typicode.com/posts", values)

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
