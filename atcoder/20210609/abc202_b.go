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

func init() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
}

func changeNumber(s string) string {
	nums := map[string]string{
		"0": "0",
		"1": "1",
		"6": "9",
		"8": "8",
		"9": "6",
	}

	return nums[s]
}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func main() {
	str := strings.Split(reverse(readline()), "")
	newStr := []rune("")
	for _, s := range str {
		newStr += changeNumber(s)
	}
	fmt.Println(newStr)
}
