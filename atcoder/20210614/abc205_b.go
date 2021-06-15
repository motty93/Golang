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

func checkDeplicate(args []int) bool {
	s := map[int]bool{}

	for i := 0; i < len(args); i++ {
		if !s[args[i]] {
			s[args[i]] = true
		} else {
			return true
		}
	}

	return false
}

func main() {
	var n int
	fmt.Scan(&n)
	an := readIntSlice()
	if checkDeplicate(an) {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
