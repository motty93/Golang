package main

import "fmt"

// Book is struct
type Book struct {
	title string
}

// 領域確保
func main() {
	bookList := []*Book{} // Book structのslice

	for i := 0; i < 10; i++ {
		book := new(Book)
		book.title = fmt.Sprintf("Title#%d", i)
		// bookを入れたものを再代入する
		bookList = append(bookList, book)
	}
	for _, book := range bookList {
		fmt.Println(book.title)
	}
}
