package main

import (
	"fmt"
	"regexp"
	"strings"
)

var s string

func main() {
	t := "oxx"
	fmt.Scan(&s)
	r := regexp.MustCompile(s)
	if r.MatchString(strings.Repeat(t, len(s)/3+5)) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
