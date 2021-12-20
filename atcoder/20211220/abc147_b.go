package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	count := 0
	for idx, v := range s {
		if idx == len(s)/2 {
			break
		}
		if idx == 0 && s[len(s)-idx-1:] != string([]rune{v}) {
			count++
		} else if s[len(s)-idx-1:len(s)-idx] != string([]rune{v}) {
			count++
		}
	}
	fmt.Println(count)
}
