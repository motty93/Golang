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

func readIntSlice() []int {
	slice := make([]int, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, stoi(v))
	}
	return slice
}

func init() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
}

func main() {
	hw := readIntSlice()
	min := math.Inf(0)
	s := make([]int, 0)
	for i := 0; i < hw[0]; i++ {
		line := readIntSlice()
		for _, a := range line {
			if min > float64(a) {
				min = float64(a)
			}
			s = append(s, a)
		}
	}

	sum := 0
	for _, v := range s {
		if v > int(min) {
			sum = sum + v - int(min)
		}
	}
	fmt.Println(sum)
}
