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
	data := readIntSlice()
	a, b := data[0], data[1]
	if a+b >= 15 && b >= 8 {
		fmt.Println(1)
	} else if a+b >= 10 && b >= 3 {
		fmt.Println(2)
	} else if a+b >= 3 {
		fmt.Println(3)
	} else {
		fmt.Println(4)
	}
}
