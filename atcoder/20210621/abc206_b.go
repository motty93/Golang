package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const BUFSIZE = 10000000

var rdr *bufio.Reader

func stoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func init() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
}

func main() {
	var n int
	fmt.Scan(&n)
	sum := 0
	day := 0
	for i := 1; sum < n; i++ {
		sum += i
		day++
	}
	fmt.Println(day)
}
