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
	s := strings.Split(readline(), "")
	str := "Good"
	for i := 0; i < len(s); i++ {
		if i == 3 {
			break
		}

		if s[i] == s[i+1] {
			str = "Bad"
		}
	}

	fmt.Println(str)
}
