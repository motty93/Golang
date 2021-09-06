package main

import (
	"bufio"
	"fmt"
	"math/big"
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

func readIntSlice() []int64 {
	slice := make([]int64, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, int64(stoi(v)))
	}
	return slice
}

func init() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
}

func main() {
	var n int
	fmt.Scan(&n)
	an := readIntSlice()
	for _, a := range an {
		if a == 0 {
			fmt.Println(0)
			return
		}
	}

	result := big.NewInt(1)
	for _, a := range an {
		result.Mul(result, big.NewInt(a))

		if result.Cmp(big.NewInt(1000000000000000000)) == 1 {
			break
		}
	}

	if result.Cmp(big.NewInt(1000000000000000000)) == 1 {
		fmt.Println(-1)
	} else {
		fmt.Println(result)
	}
}
