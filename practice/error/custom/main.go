package main

import (
	"fmt"
	"time"
)

type (
	Article struct {
		postedAt time.Time
	}
	// カスタムエラー構造体
	ArticleNotFoundError struct {
		Message string
	}
)

func main() {
	article, err := GetArticle()
	if err != nil {
		fmt.Println(err.Error()) // => "article not found"
	}

	fmt.Println(article)
}

func GetArticle() (*Article, error) {
	return nil, NewArticleNotFoundError("article not found")
}

func NewArticleNotFoundError(message string) *ArticleNotFoundError {
	return &ArticleNotFoundError{Message: message}
}

// このメソッドがあるおかげで、ArticleNotFoundError構造体の型はerrorになり、GetArticleの第2引数の型はerrorになりました。
func (e *ArticleNotFoundError) Error() string {
	return e.Message
}
