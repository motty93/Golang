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

func main() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
	n, an := stoi(readline()), readIntSlice()
	alice, bob := 0, 0
	sort.SliceStable(an, func(i, j int) bool { return an[i] > an[j] })
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			alice += an[i]
		} else {
			bob += an[i]
		}
	}
	fmt.Println(alice - bob)
}
