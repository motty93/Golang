package main

import (
	"fmt"
	"strconv"
	"strings"
)

func stoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func main() {
	var s1, s2, s3, t string
	fmt.Scan(&s1, &s2, &s3, &t)
	ts := strings.Split(t, "")
	sn := []string{s1, s2, s3}
	r := ""
	for _, v := range ts {
		r += sn[stoi(v)-1]
	}
	fmt.Println(r)
}
