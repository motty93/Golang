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

func main() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
	var n int
	fmt.Scan(&n)
	an := readIntSlice()
	bn := readIntSlice()
	result := make([]int, 0)
	for idx, num := range an {
		for i := num; bn[idx] < i; i++ {
			result = append(result, i)
		}
	}
	fmt.Println(result)
}
