package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const BUFSIZE = 10000000

var rdr *bufio.Reader

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
	var n int
	fmt.Scan(&n)
	st := readSlice()
	s := strings.Split(st[0], "")
	t := strings.Split(st[1], "")
	newStr := ""
	for i := 0; i < n; i++ {
		newStr = newStr + s[i] + t[i]
	}
	fmt.Println(newStr)
}
