package main

import (
	"fmt"

	"./mylib"
	"./mylib/under"
)

// ファイル名は影響しない
func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))
	mylib.Say()
	under.Hello()
	person := mylib.Person{"mike", 20}
	fmt.Println(person)
	fmt.Println(mylib.Public)
	// lower caseだとアクセスできない
	// fmt.Println(mylib.private)
}
