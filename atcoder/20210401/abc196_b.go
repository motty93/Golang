package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Scan(&x)
	s := strings.Split(x, ".")
	fmt.Println(s[0])
}
