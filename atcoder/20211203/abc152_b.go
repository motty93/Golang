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

func itos(i int) string {
	return strconv.Itoa(i)
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

func readStringSlice() []string {
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
	ab := readStringSlice()
	if ab[0] >= ab[1] {
		fmt.Println(strings.Repeat(ab[1], stoi(ab[0])))
	} else {
		fmt.Println(strings.Repeat(ab[0], stoi(ab[1])))
	}
}
