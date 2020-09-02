package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Article struct {
	Title string
	Body  string
}

var tpl *template.Template

func init() {
	funcMap := template.FuncMap{
		"safehtml": func(text string) template.HTML { return template.HTML(text) },
	}
	tpl = template.Must(template.New("").Funcs(funcMap).ParseFiles("template.html"))
}

func helloHandler(rw http.ResponseWriter, req *http.Request) {
	article := Article{
		Title: "golang practice",
		Body:  "<h1>hello golang</h1>",
	}

	if err := tpl.ExecuteTemplate(rw, "template.html", article); err != nil {
		log.Fatalln(err.Error())
	}
}

func dogHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "<h1>dogs</h1><h2>dog dog dog</h2>")
}

func main() {
	// http.Handle("/hello", http.HandlerFunc(helloHandler))
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/dog", dogHandler)

	http.ListenAndServe(":3000", nil)
}
