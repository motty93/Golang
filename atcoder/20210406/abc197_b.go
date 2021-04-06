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

func readStringSlice() []string {
	slice := make([]string, 0)
	lines := strings.Split(readline(), "")
	for _, v := range lines {
		slice = append(slice, v)
	}
	return slice
}

func main() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
	hwxy := readIntSlice()
	sh := make([]string, hwxy[0])
	for i := 0; i < hwxy[0]; i++ {
		sh[i] = readStringSlice()
	}
	fmt.Println(sh)
	fmt.Println(sh[0])
}
