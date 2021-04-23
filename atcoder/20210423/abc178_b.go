package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/adam-lavrik/go-imath/ix"
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

func main() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
	row := readIntSlice()
	ab := []int{row[0], row[1]}
	cd := []int{row[2], row[3]}
	slices := make([]int, 0)
	for _, i := range ab {
		for _, j := range cd {
			slices = append(slices, i*j)
		}
	}
	fmt.Println(ix.MaxSlice(slices))
}
