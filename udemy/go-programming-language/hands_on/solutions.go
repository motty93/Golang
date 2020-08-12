package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.Handle("/", http.HandlerFunc(Root))
	http.Handle("/dog/", http.HandlerFunc(MyDog))
	http.HandleFunc("/me/", MyName)
	http.ListenAndServe(":8080", nil)
}

// Root is app root
func Root(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "this is app root")
}

// MyDog is dog path
func MyDog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog dog dog")
}

// MyName is me path
func MyName(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("something.gohtml")
	if err != nil {
		// , errをつけるとあとに具体的なエラー内容が入る
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(w, "something.gohtml", "test")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
