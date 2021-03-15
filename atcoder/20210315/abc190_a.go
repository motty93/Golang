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
	s := make([]int, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		s = append(s, stoi(v))
	}
	return s
}

func main() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
	nums := readIntSlice()
	if nums[0] > nums[1] {
		fmt.Println("Takahashi")
	} else if nums[0] == nums[1] {
		if nums[2] == 1 {
			fmt.Println("Takahashi")
		} else {
			fmt.Println("Aoki")
		}
	} else {
		fmt.Println("Aoki")
	}
}
