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
	xy := readIntSlice()
	if xy[1]%2 != 0 {
		fmt.Println("No")
		return
	}

	b := (xy[1] - 2*xy[0]) / 2
	if xy[0]-b < 0 || b < 0 {
		fmt.Println("No")
		return
	}

	fmt.Println("Yes")
}
