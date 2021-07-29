package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	ln := readIntSlice()

	sort.SliceStable(ln, func(i, j int) bool { return ln[i] < ln[j] })
	sum := 0
	for i := 0; i < len(ln); i++ {
		for j := i + 1; j < len(ln); j++ {
			for k := j + 1; k < len(ln); k++ {
				if ln[i]+ln[j] > ln[k] && ln[i] != ln[j] && ln[j] != ln[k] {
					sum++
				}
			}
		}
	}

	fmt.Println(sum)
}
