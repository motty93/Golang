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
	hw := readIntSlice()
	a := [][]int{}
	for i := 0; i < hw[0]; i++ {
		a = append(a, readIntSlice())
	}

	for i := 0; i < hw[0]-1; i++ {
		for j := 0; j < hw[1]-1; j++ {
			if a[i][j]+a[i+1][j+1] > a[i+1][j]+a[i][j+1] {
				fmt.Println("No")
				os.Exit(0)
			}
		}
	}

	fmt.Println("Yes")
}
