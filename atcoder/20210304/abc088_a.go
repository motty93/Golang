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
	n := stoi(readline())
	a := stoi(readline())
	if n%500 > a {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
