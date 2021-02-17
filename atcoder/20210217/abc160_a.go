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
	row := strings.Split(readline(), "")
	if row[2] == row[3] && row[4] == row[5] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
