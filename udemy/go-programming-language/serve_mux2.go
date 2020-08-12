package main

import (
	"io"
	"net/http"
)

type hotdog int
type hotcat int

// func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
// 	io.WriteString(res, "doggy doggy doggy")
// }
//
// func (m hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
// 	io.WriteString(res, "cat cat cat")
// }

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy doggy doggy")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}
func main() {
	// var d hotdog
	// var c hotcat

	// mux := http.NewServeMux()
	// mux.Handle("/dog/", d)
	// mux.Handle("/cat/", c)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat/", c)
	// path: /cat/something/another
	http.ListenAndServe(":8080", nil)
}
