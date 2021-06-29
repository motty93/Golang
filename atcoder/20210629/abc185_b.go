package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BUFSIZE = 10000000

var rdr *bufio.Reader

func stoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

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

func readIntSlice() []int {
	slice := make([]int, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, stoi(v))
	}
	return slice
}

func init() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
}

func main() {
	nmt := readIntSlice()
	ab := make([][]int, nmt[1])
	ans := nmt[0]
	if len(ab) > 1 {
		for i := 0; i < nmt[1]; i++ {
			ab[i] = readIntSlice()

			switch i {
			case 0:
				ans -= ab[i][0]
				if ans <= 0 {
					fmt.Println("No")
					return
				}
				ans += ab[i][1] - ab[i][0]
				if ans > nmt[0] {
					ans = nmt[0]
				}
			case nmt[1] - 1:
				ans -= ab[i][0] - ab[i-1][1]
				if ans <= 0 {
					fmt.Println("No")
					return
				}
				ans += ab[i][1] - ab[i][0]
				if ans > nmt[0] {
					ans = nmt[0]
				}
				ans -= nmt[2] - ab[i][1]
			default:
				ans -= ab[i][0] - ab[i-1][1]
				if ans <= 0 {
					fmt.Println("No")
					return
				}
				ans += ab[i][1] - ab[i][0]
				if ans > nmt[0] {
					ans = nmt[0]
				}
			}
		}
	} else {
		ab[0] = readIntSlice()
		ans -= ab[0][0]
		if ans <= 0 {
			fmt.Println("No")
			return
		}
		ans += ab[0][1] - ab[0][0]
		if ans > nmt[0] {
			ans = nmt[0]
		}
		ans -= nmt[2] - ab[0][1]
		if ans <= 0 {
			fmt.Println("No")
			return
		}
	}

	if ans <= 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
