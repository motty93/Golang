package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Article struct {
	Title string
	Body  string
}

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

var tpl *template.Template

func safeHTMLTemplate(text string) template.HTML {
	return template.HTML(text)
}

func init() {
	funcMap := template.FuncMap{
		"safehtml": safeHTMLTemplate,
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

func jsonResponseHandler(rw http.ResponseWriter, req *http.Request) {
	data := map[string]interface{}{
		"message": "hello golang",
		"status":  http.StatusOK,
	}

	// map to json
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(rw, string(bytes))
}

func productResponseHandler(rw http.ResponseWriter, req *http.Request) {
	// product := Product{"商品A", 100, 10}
	product := Product{
		Name:     "商品A",
		Price:    100,
		Quantity: 10,
	}

	// struct to json byte
	bytes, err := json.Marshal(product)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(bytes))
}

func main() {
	// http.Handle("/hello", http.HandlerFunc(helloHandler))
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/dog", dogHandler)
	http.HandleFunc("/json", jsonResponseHandler)
	http.HandleFunc("/product", productResponseHandler)

	http.ListenAndServe(":3000", nil)
}
