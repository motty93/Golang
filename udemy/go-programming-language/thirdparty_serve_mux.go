package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/", index)
	// mux.GET("/about", about)
	// mux.GET("/contact", contact)
	// mux.GET("/apply", apply)
	// mux.POST("/apply", applyProcess)
	// mux.GET("/user/:name/exact", user)
	// mux.GET("/blog/:category/:article", blogRead)
	// mux.POST("/blog/:category/:article", blogWrite)
	http.ListenAndServe(":8080", mux)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

// HandleError is http response error function
func HandleError(w http.ResponseWriter, e error) {
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		log.Println(e)
	}
}
