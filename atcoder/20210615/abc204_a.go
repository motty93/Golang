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

func init() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
}

func remove(strings *[]string, target string) {
	for _, v := range *strings {
		if v != target {
			*strings = append(*strings, v)
		}
	}
}

func main() {
	var x, y string
	xyz := []string{"0", "1", "2"}
	fmt.Scan(&x, &y)
	remove(&xyz, x)
	remove(&xyz, y)

	fmt.Println(len(xyz))
	fmt.Println(xyz)
	if len(xyz) > 1 {
		fmt.Println(x)
	} else {
		fmt.Println(xyz[0])
	}
}
