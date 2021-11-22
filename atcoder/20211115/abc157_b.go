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

func main() {
	an := map[int]bool{}
	for i := 0; i < 3; i++ {
		d := readIntSlice()
		for _, v := range d {
			an[v] = true
		}
	}
	fmt.Println(an)
	n := stoi(readline())
	for i := 0; i < n; i++ {
		num := stoi(readline())
		if an[num] {
			an[num] = false
		}
	}
	fmt.Println(an)
}
