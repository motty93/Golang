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

// 1行をstringで読み込み
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

var dn int

func main() {
	nk := readIntSlice()
	var adk [][]int
	for i := 0; i < nk[1]; i++ {
		dn = readIntSlice()[0]
		adk = append(adk, readIntSlice())
	}

	uniq := []int{}
	m := make(map[string]bool)
	for _, a := range adk {
		for _, v := range a {
			if !m[strconv.Itoa(v)] {
				m[strconv.Itoa(v)] = true
				uniq = append(uniq, v)
			}
		}
	}

	fmt.Println(nk[0] - len(uniq))
}
