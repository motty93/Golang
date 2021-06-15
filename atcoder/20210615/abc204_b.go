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
var n int

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

func readIntSlice() []int {
	s := make([]int, 0)
	l := strings.Split(readline(), " ")
	for _, v := range l {
		s = append(s, stoi(v))
	}
	return s
}

func init() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
}

func main() {
	fmt.Scan(&n)
	an := readIntSlice()
	sum := 0
	for _, a := range an {
		if a <= 10 {
			continue
		}

		sum += a - 10
	}
	fmt.Println(sum)
}
