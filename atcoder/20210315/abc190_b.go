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
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		s = append(s, stoi(v))
	}
	return s
}

func main() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
	nsd := readIntSlice()
	result := false

	for i := 0; i < nsd[0]; i++ {
		xy := readIntSlice()

		if xy[0] < nsd[1] && xy[1] > nsd[2] {
			result = true
		}
	}

	if result {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
