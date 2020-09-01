package main

import (
	"fmt"
	"net/http"
)

// 使わないので一旦コメントアウト
// func ServeHTTP(rw http.ResponseWriter, req *http.Request) {
// 	switch req.URL.Path {
// 	case "/dog":
// 		fmt.Fprintf(rw, "dog dog dog")
// 	case "/cat":
// 		fmt.Fprintf(rw, "cat cat cat")
// 	default:
// 		fmt.Fprintf(rw, "hello world")
// 	}
// }

func helloHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, `
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>golang practice</title>
			</head>
			<body>
				<h1>hello golang</h1>
			</body>
		</html>
	`)
}

func dogHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, `
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>golang practice</title>
			</head>
			<body>
				<h1>dogs</h1>
				<h2>dog dog dog</h2>
			</body>
		</html>
	`)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/dog", dogHandler)

	http.ListenAndServe(":8080", nil)
}
