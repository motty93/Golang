package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const BUFSIZE = 10000000

var rdr *bufio.Reader

// 1行をstringで読み込み
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
	var n int
	fmt.Scan(&n)
	names := make([]string, 0)
	for i := 0; i < n; i++ {
		names = append(names, strings.Join(readSlice(), ""))
	}

	exist := false
	for j := 0; j < len(names); j++ {
		for k := j + 1; k < len(names); k++ {
			if names[j] == names[k] {
				exist = true
				break
			}
		}
	}

	if exist {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
