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
	b := make([]byte, 0, 16)
	for {
		l, p, _ := rdr.ReadLine()
		b = append(b, l...)
		if !p {
			break
		}
	}

	return string(b)
}

func readIntSlice() []int {
	s := make([]int, 0)
	l := strings.Split(readline(), " ")
	for _, v := range l {
		s = append(s, stoi(v))
	}
	return s
}

func init() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
}

func main() {
	nk := readIntSlice()
	pn := readIntSlice()
	sort.Slice(pn, func(i, j int) bool { return pn[i] < pn[j] })
	sum := 0
	for i := 0; i < nk[1]; i++ {
		sum += pn[i]
	}
	fmt.Println(sum)
}
