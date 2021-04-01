package main

import (
	"bufio"
	"fmt"
	"os"
)

const BUFSIZE = 10000000

var rdr *bufio.Reader

func main() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
	var a, b, w int
	fmt.Scanf("%d %d %d", &a, &b, &w)
	w *= 1000
}
