package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const BUFSIZE = 10000000

var rdr *bufio.Reader

func stoi(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}

// 1行をstringで読み込み
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

func main() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
	j := stoi(readline())
	k := j / 500
	h := (j % 500) / 5
	fmt.Printf("%d", k*1000+h*5)
}
