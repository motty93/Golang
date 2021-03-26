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

func relu(x int) int {
	if x >= 0 {
		return x
	} else {
		return 0
	}
}

func main() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
	x := stoi(readline())
	fmt.Println(relu(x))
}
