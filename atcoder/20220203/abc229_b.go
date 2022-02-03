package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BUFSIZE = 10000000

var (
	rdr *bufio.Reader
)

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

func readSlice() []string {
	slice := make([]string, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, v)
	}
	return slice
}

func init() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
}

func main() {
	ab := readSlice()
	an := strings.Split(ab[0], "")
	bn := strings.Split(ab[1], "")

	if len(an) >= len(bn) {
		for i := len(bn) - 1; i >= 0; i-- {
			if stoi(bn[i])+stoi(an[i+1]) > 9 {
				fmt.Println("Hard")
				os.Exit(0)
			}
		}
	} else {
		for i := len(an) - 1; i >= 0; i-- {
			if stoi(bn[i+1])+stoi(an[i]) > 9 {
				fmt.Println("Hard")
				os.Exit(0)
			}
		}
	}

	fmt.Println("Easy")
}
