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
	d := make(map[int]bool)
	num := 0
	u := []int{}
	for i := 0; i < n; i++ {
		num = stoi(readline())
		if !d[num] {
			d[num] = true
			u = append(u, num)
		}
	}
	fmt.Println(len(u))
}
