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
	buf := make([]byte, 0, 16)
	for {
		l, p, _ := rdr.ReadLine()
		buf = append(buf, l...)
		if !p {
			break
		}
	}

	return string(buf)
}

func main() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
	x := stoi(readline())
	a := stoi(readline())
	b := stoi(readline())

	x -= a
	for {
		if x < b {
			break
		}

		x -= b
	}
	fmt.Println(x)
}
