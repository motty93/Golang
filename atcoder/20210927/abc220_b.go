package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		line, prefix, _ := rdr.ReadLine()
		buf = append(buf, line...)
		if !prefix {
			break
		}
	}

	return string(buf)
}

func readSlice() []string {
	slice := make([]string, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, v)
	}
	return slice
}

func init() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
}

func main() {
	var k int
	fmt.Scan(&k)
	ab := readSlice()
	a := 0
	b := 0
	for idx, v := range strings.Split(ab[0], "") {
		fmt.Println(stoi(v))
	}
}
