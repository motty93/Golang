package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const BUFSIZE = 10000000

var rdr *bufio.Reader

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
	r := strings.Split(readline(), " ")
	fmt.Printf("%s %s %s", r[2], r[0], r[1])
}
