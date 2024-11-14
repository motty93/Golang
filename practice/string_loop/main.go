package main

import "fmt"

func main() {
	str := "Hello, World!"
	strSlice := make([]string, 0)
	for _, r := range str {
		strSlice = append(strSlice, string(r))
	}

	fmt.Println(len(strSlice))
	fmt.Println(len(str))
}
