package main

import (
	"fmt"
	"regexp"
)

func main() {
	var slot string
	fmt.Scan(&slot)
	r := regexp.MustCompile(`[A-Z]{3}`)

	if r.MatchString(slot) {
		fmt.Println("Won")
	} else {
		fmt.Println("Lost")
	}

