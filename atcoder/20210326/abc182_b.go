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

func stoUint(s string) uint64 {
	ui, _ := strconv.ParseUint(s, 10, 64)
	return ui
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

func gcd(m, n uint64) uint64 {
	x := new(big.Int)
	y := new(big.Int)
	z := new(big.Int)
	a := new(big.Int).SetUint64(m)
	b := new(big.Int).SetUint64(n)
	return z.GCD(x, y, a, b).Uint64()
}

func main() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
	var n int
	fmt.Scan(&n)
	an := readIntSlice()
	sum := 0
	gcdNum := []int{}
	max := -999999

	// GCD度の計算
	for _, a := range an {
		for _, b := range an {
			if b%a == 0 {
				sum += 1
			}
		}
		gcdNum = append(gcdNum, sum)
		sum = 0
	}

	fmt.Println(gcdNum)
	// GCD度最大値取得
	for _, s := range gcdNum {
		if max < s {
			max = s
		}
	}

	for idx, a := range gcdNum {
		if a == max {
			fmt.Println(an[idx])
			break
		}
	}
}
