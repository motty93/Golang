package main

import "fmt"

func remove(strs []string, search string) []string {
	r := make([]string, 0)
	for _, v := range strs {
		if v != search {
			r = append(r, v)
		}
	}

	return r
}

func main() {
	var a, b, c string
	fmt.Scan(&a, &b, &c)
	sn := []string{"ABC", "ARC", "AGC", "AHC"}
	sn = remove(sn, a)
	sn = remove(sn, b)
	sn = remove(sn, c)
	fmt.Println(sn[0])
}
