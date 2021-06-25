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
		line, prefix, _ := rdr.ReadLine()
		buf = append(buf, line...)
		if !prefix {
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
	xyn := make([][]int, 0)
	for i := 0; i < n; i++ {
		xyn = append(xyn, readIntSlice())
	}

	cnt := 0
	for i, v1 := range xyn {
		for j, v2 := range xyn {
			if i >= j {
				continue
			}

			if math.Abs(float64(v2[1]-v1[1])) <= math.Abs(float64(v2[0]-v1[0])) {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}
