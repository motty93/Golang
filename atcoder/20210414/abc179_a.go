package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	if strings.HasSuffix(s, "s") {
		fmt.Println(s + "es")
	} else {
		fmt.Println(s + "s")
	}
}
