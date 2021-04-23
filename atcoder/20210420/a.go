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
	slice := make([]int, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, stoi(v))
	}
	return slice
}

func main() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
	row := readIntSlice()
	eat := row[1]
	last := row[1] + row[0] - 1

	if last <= 0 {
		eat = last
	} else if row[1] > 0 {
		eat = row[1]
	} else {
		eat = 0
	}

	result := ((last-row[1]+1)*(row[1]+last)/2 - eat)
	fmt.Println(result)
}
