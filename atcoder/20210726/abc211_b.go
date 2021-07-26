package main

import "fmt"

func remove(strs *[]string, target string) {
	r := []string{}
	for _, s := range *strs {
		if s != target {
			r = append(r, s)
		}
	}
	*strs = r
}

func main() {
	var s1, s2, s3, s4 string
	fmt.Scan(&s1, &s2, &s3, &s4)
	h := []string{"H", "2B", "3B", "HR"}
	remove(&h, s1)
	remove(&h, s2)
	remove(&h, s3)
	remove(&h, s4)
	if len(h) == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
