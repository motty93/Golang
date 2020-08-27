package main

import (
	"fmt"
	"net/http"
)

// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

type MywebserverType bool

func (m MywebserverType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>golang</title>
		</head>
		<body>
			<h1>Golang HTTP Test</h1>
		</body>
		</html>
	`)
}

func main() {
	var k MywebserverType
	http.ListenAndServe("localhost:8080", k)
}
