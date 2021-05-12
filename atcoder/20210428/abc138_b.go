package main

import (
	"bufio"
	"fmt"
	"math/big"
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
	ai := readIntSlice()
	sum := new(big.Rat)
	for _, a := range ai {
		sum = new(big.Rat).Add(sum, big.NewRat(1, int64(a)))
	}
	sum, _ = sum.Float64
	fmt.Println(big.NewRat(1, sum))
}
