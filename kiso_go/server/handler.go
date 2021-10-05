package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{Id: 1, Name: "motty", Age: 30},
	{Id: 2, Name: "motty2", Age: 20},
	{Id: 3, Name: "motty child", Age: 2},
}

func findByUser(id int) *User {
	for _, u := range users {
		if u.Id == id {
			return &u
		}
	}

	return nil
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "<html><body>hello world</body></html>")
}

func stoi(s string) int {
	result, ok := strconv.Atoi(s)
	if ok != nil {
		panic("Failed: " + s + " can't convert to int.")
	}
	return result
}

func getUserHandler(rw http.ResponseWriter, req *http.Request) {
	sub := strings.TrimPrefix(req.URL.Path, "/users")
	_, id := filepath.Split(sub)
	user := findByUser(stoi(id))

	bytes, err := json.Marshal(&user)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(bytes))
}

func getUsersHandler(rw http.ResponseWriter, req *http.Request) {
	bytes, err := json.Marshal(&users)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(bytes))
}
