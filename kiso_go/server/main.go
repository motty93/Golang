package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, req *http.Request) {
		val := req.FormValue("say")

		fmt.Fprintf(w, "<html><body>echo: %s</body></html>", html.EscapeString(val))
	})
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/users/", getUserHandler)
	http.HandleFunc("/users", getUsersHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
