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

func itos(i int) string {
	return strconv.Itoa(i)
}

func main() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
	nx := readIntSlice()
	ai := readIntSlice()
	r := []string{}
	for _, a := range ai {
		if a != nx[1] {
			r = append(r, itos(a))
		}
	}
	fmt.Println(strings.Join(r, " "))
}
