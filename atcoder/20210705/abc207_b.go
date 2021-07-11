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

func stoi(s string) int64 {
	r, _ := strconv.ParseInt(s, 10, 64)
	return r
}

func readline() string {
	buf := make([]byte, 0, 32)
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
	slice := make([]int64, 0)
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
	n := readIntSlice()
	var i int64 = -1
	if (n[2]*n[3] - n[1]) < 0 {
		fmt.Println(i)
	} else {
		for {
			i++
			if n[0] <= (n[2]*n[3]-n[1])*i {
				break
			}
		}
		fmt.Println(i)
	}
}
