package main

import (
	"fmt"
	"io"
	"net/http"
)

// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

func ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "dog dog dog")
	case "/":
		fmt.Fprintf(w, htmlTemplate("root", "<h1>Hello Golang</h1>"))
	}
}

func login(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, htmlTemplate("login", "<h1>enter your username and password</h1>"))
}

func welcome(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, htmlTemplate("welcome", "<h1>welcome</h1>"))
}

func htmlTemplate(title, body string) (template string) {
	template = fmt.Sprintf(
		`<!DOCTYPE html>
			<html lang="en">
				<head>
					<meta charset="UTF-8">
					<title>%s</title>
					%s
				</head>
				<body>
					%s
				</body>
			</html>`, title, title, body)

	return
}

func main() {
	// http.Handle("/", http.HandlerFunc(login))と同じ
	http.HandleFunc("/", login)
	http.HandleFunc("/login", login)
	// http.HandleFunc("/welcome", welcome)
	http.HandleFunc("/welcome/", welcome) //最後にスラッシュを入れると/welcome/testでも普通に表示される

	// http.ListenAndServe("localhost:8080", http.HandlerFunc(ServeHTTP))
	http.ListenAndServe("localhost:8080", nil) //nilは何もなければ404を返す
}
