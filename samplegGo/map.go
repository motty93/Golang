package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["i"] = 23
	m["j"] = -1
	m["k"] = 0
	fmt.Println("map:", m)
	fmt.Println("len:", len(m))

	v1 := m["i"]
	fmt.Println("v1:", v1)

	delete(m, "k")
	fmt.Println("del:", m)

	// ２つ目の返り値は真偽値
	i, prs := m["i"]
	fmt.Println("prs:", prs)
	fmt.Println("i:", i)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n:", n)
}
