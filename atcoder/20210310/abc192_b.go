package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
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
	strs := strings.Split(readline(), "")
	even := true
	odd := true

	for idx, s := range strs {
		if idx%2 == 0 {
			even = even && unicode.IsLower(rune(s[0]))
		} else {
			odd = odd && unicode.IsUpper(rune(s[0]))
		}
	}

	if even && odd {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
