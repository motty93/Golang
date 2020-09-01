package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Article struct {
	Title string
	Body  template.HTML
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template.html"))
}

func escapeHTML(html string) (tmpl template.HTML) {
	tmpl = template.HTML(html)

	return
}

func helloHandler(rw http.ResponseWriter, req *http.Request) {
	article := Article{
		Title: "golang practice",
		Body:  escapeHTML("<h1>hello golang</h1>"),
	}

	if err := tpl.ExecuteTemplate(rw, "template.html", article); err != nil {
		log.Fatalln(err.Error())
	}
}

func dogHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "<h1>dogs</h1><h2>dog dog dog</h2>")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/dog", dogHandler)

	http.ListenAndServe(":3000", nil)
}
