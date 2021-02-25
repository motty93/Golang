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

func readline() string {
	b := make([]byte, 0, 16)
	for {
		l, p, _ := rdr.ReadLine()
		b = append(b, l...)
		if !p {
			break
		}
	}

	return string(b)
}

func main() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
	n := stoi(readline())
	sum := 0
	for i := 1; i <= n; i++ {
		if i%3 != 0 && i%5 != 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}
