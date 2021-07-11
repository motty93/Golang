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
	var n int
	fmt.Scan(&n)
	xn := readIntSlice()
	max := -999999.0
	sum := 0.0
	squaredSum := 0.0
	for _, v := range xn {
		x := math.Abs(float64(v))
		if max < x {
			max = x
		}
		sum += x
		squaredSum += math.Pow(x, 2.0)
	}
	fmt.Printf("%d\n%g\n%d\n", int(sum), math.Sqrt(squaredSum), int(max))
}
