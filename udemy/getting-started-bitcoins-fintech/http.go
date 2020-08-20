package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// 単純なGET
	// resp, _ := http.Get("http://example.com")
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	// header情報など入れたい時のrequest
	base, _ := url.Parse("http://example.com/asdfasdfsa")
	reference, _ := url.Parse("/test?a=1&b=10")
	// baseが間違っていても修正してくれる
	endpoint := base.ResolveReference(reference).String()
	// resp, _ := http.Get(endpoint)
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(body)
	fmt.Println(endpoint)

	req, _ := http.NewRequest("GET", endpoint, nil)
	// POST
	// req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("password")))
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	q := req.URL.Query()
	// クエリ追加
	q.Add("c", "3&%")
	fmt.Println(q)
	// &などを渡したい場合Encodeして渡し、受取先でDecodeする
	fmt.Println(q.Encode())
	req.URL.RawQuery = q.Encode()

	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
