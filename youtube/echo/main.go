package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseData struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Response int    `json:"status"`
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, htmlTemplate("hello golang", "<h1>hellor world</h1>"))
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Fprintf(w, htmlTemplate("GET", "<h1>Using GET for login endpoint</h1>"))
	case "POST":
		fmt.Fprintf(w, htmlTemplate("POST", "<h1>Using POST for login endpoint</h1>"))
	default:
		fmt.Fprintf(w, htmlTemplate("login", "<h1>login endpoint</h1>"))
	}
}

func welcomeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, htmlTemplate("welcome", "<h1>welcome</h1>"))
}

func responseHandler(w http.ResponseWriter, req *http.Request) {
	data := ResponseData{"motty", 40, http.StatusOK}
	res, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(res))
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

type test int

func (l test) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "test")
}

func main() {
	// http.Handle("/", http.HandlerFunc(login))と同じ
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/response", responseHandler)
	// http.HandleFunc("/welcome", welcome)
	http.HandleFunc("/welcome/", welcomeHandler) //最後にスラッシュを入れると/welcome/testでも普通に表示される

	// interfaceを使ったServeHTTPのオーバーライド
	var t test
	http.Handle("/test", t)

	// http.ListenAndServe("localhost:8080", http.HandlerFunc(ServeHTTP))
	http.ListenAndServe("localhost:8080", nil) //nilは何もなければ404を返す
}
