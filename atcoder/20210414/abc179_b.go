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
	r := 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		d := readIntSlice()

		if r == 3 {
			break
		} else if d[0] == d[1] {
			r += 1
		} else {
			r = 0
		}
	}

	if r == 3 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
