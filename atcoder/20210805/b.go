package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const BUFSIZE = 10000000

var rdr *bufio.Reader

func stof(s string) float64 {
	r, _ := strconv.Atoi(s)
	return float64(r)
}

// 1行をstringで読み込み
func readline() string {
	b := make([]byte, 0, 16)
	for {
		l, p, _ := rdr.ReadLine()
		b = append(b, l...)
		if !p {
			break
		}
	}

	return string(b)
}

func readIntSlice() []float64 {
	s := make([]float64, 0)
	l := strings.Split(readline(), " ")
	for _, v := range l {
		s = append(s, stof(v))
	}
	return s
}

func init() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
}

func main() {
	nd := readIntSlice()
	d2 := math.Pow(nd[1], 2.0)
	count := 0
	for i := 0; i < int(nd[0]); i++ {
		xy := readIntSlice()
		if math.Pow(xy[0], 2.0)+math.Pow(xy[1], 2.0) <= d2 {
			count++
		}
	}
	fmt.Println(count)
}
