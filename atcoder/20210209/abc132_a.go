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
	m := make(map[string]bool)
	uniq := []string{}

	for _, ele := range s {
		if !m[ele] {
			m[ele] = true
			uniq = append(uniq, ele)
		}
	}
	if len(uniq) != 2 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
