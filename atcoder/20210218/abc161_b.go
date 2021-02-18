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
	nm := readIntSlice()
	ai := readIntSlice()
	sum := 0
	for _, v := range ai {
		sum += v
	}

	tmp := 0
	for _, v := range ai {
		if v >= sum/(4*nm[1]) {
			tmp += 1
		}
	}

	if tmp >= nm[1] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
