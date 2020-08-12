package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

// cd hands_on; go run solutions.go
func main() {
	http.HandleFunc("/", Root)
	http.HandleFunc("/dog/", MyDog)
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
		log.Fatalln("error parsing template", err)
	}
	err = tpl.ExecuteTemplate(w, "something.gohtml", nil)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
