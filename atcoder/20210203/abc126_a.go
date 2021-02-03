package main

import (
	"fmt"
	"strconv"
	"strings"
)

func stringToInt(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}

var (
	s, n, k, result string
	slice           []string
)

func main() {
	fmt.Scan(&n, &k, &s)
	K := stringToInt(k)
	slice = strings.Split(s, "")
	slice[K-1] = strings.ToLower(slice[K-1])
	fmt.Printf("%s\n", strings.Join(slice, ""))
}
