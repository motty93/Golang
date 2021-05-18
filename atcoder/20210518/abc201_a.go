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

func init() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
}

func main() {
	an := readIntSlice()
	sort.Slice(an, func(i, j int) bool { return an[i] < an[j] })
	if an[2]-an[1] == an[1]-an[0] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
