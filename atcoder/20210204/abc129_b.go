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
	result, ok := strconv.Atoi(s)
	if ok != nil {
		panic("Failed: " + s + " can't convert to int.")
	}
	return result
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func readline() string {
	buf := make([]byte, 0, 16)
	for {
		line, prefix, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
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

func main() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
	row := readIntSlice()
	fmt.Println(row)
}
