package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	i := []int{3, 6, 2, 1, 7, 0}
	s := []string{"d", "a", "f"}
	p := []Person{
		{"nancy", 20},
		{"vera", 40},
		{"mike", 30},
		{"bob", 10},
	}
	fmt.Println(i, s, p)

	sort.Ints(i)
	fmt.Println(i)
	sort.Strings(s)
	fmt.Println(s)
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
	fmt.Println(p)
	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })
	fmt.Println(p)
}
