package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	rep := regexp.MustCompile(`\d{4}`)
	fmt.Println(rep.ReplaceAllString(readline(), "2018"))
}
