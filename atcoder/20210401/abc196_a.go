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
	ab := readIntSlice()
	cd := readIntSlice()
	xy := -200

	for i := ab[0]; i <= ab[1]; i++ {
		for j := cd[0]; j <= cd[1]; j++ {
			if i-j >= xy {
				xy = i - j
			}
		}
	}

	fmt.Println(xy)
}
