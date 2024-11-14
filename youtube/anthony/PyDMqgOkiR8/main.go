package main

import (
	"fmt"
	"time"
)

type Book struct {
	title    string
	author   string
	numPages int

	isSaved bool
	savedAt time.Time
}

func (b *Book) save() {
	b.isSaved = true
	b.savedAt = time.Now()
}

func saveBook(b *Book) {
	b.title = "The Hobbit"
	b.author = "J.R.R. Tolkien"
	b.numPages = 310
	b.isSaved = true
	b.savedAt = time.Now()
}

func main() {
	book := Book{
		title:    "The Hobbit",
		author:   "J.R.R. Tolkien",
		numPages: 310,
	}

	book.save()
	fmt.Println("Book saved at:", book.savedAt)

	// 以下2つは一般的ではない
	book2 := Book{}
	saveBook(&book2)
	fmt.Println("Book2 saved at:", book2.savedAt)

	book3 := &Book{}
	saveBook(book3)
	fmt.Println("NoSaveBook saved at:", book3.savedAt)
}
