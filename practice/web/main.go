package main

import (
	"fmt"
	"net/http"
)

type handlerType int

func (t handlerType) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		fmt.Fprintf(rw, "dog dog dog")
	case "/cat":
		fmt.Fprintf(rw, "cat cat cat")
	default:
		fmt.Fprintf(rw, "hello world")
	}
}

func main() {
	var t handlerType
	http.ListenAndServe(":8080", t)
}
