package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(strings.Count(str, "1"))
}
