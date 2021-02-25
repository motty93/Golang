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
		l, p, _ := rdr.ReadLine()
		buf = append(buf, l...)
		if !p {
			break
		}
	}

	return string(buf)
}

func main() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
	if strings.Contains(readline(), "7") {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
