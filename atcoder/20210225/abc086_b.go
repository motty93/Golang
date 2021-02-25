package main

import (
	"bufio"
	"fmt"
	"math"
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

func isInteger(n float64) bool {
	return math.Floor(n) == n
}

func main() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
	r := strings.Fields(readline())
	n := stoi(strings.Join(r, ""))
	if isInteger(math.Sqrt(float64(n))) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
