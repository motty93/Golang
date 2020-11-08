package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("input? ")
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		if err := stdin.Err(); err != nil {
			fmt.Println(os.Stderr, err)
			break
		}
		if stdin.Text() == "exit" {
			break
		}
		fmt.Println("input is", stdin.Text())
	}
}
