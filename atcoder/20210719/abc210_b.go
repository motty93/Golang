package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, s string
	fmt.Scan(&n, &s)
	if strings.Index(s, "1")%2 == 0 {
		fmt.Println("Takahashi")
	} else {
		fmt.Println("Aoki")
	}
}
