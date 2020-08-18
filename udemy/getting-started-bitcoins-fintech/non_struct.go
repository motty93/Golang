package main

import "fmt"

// MyInt is non-struct type
type MyInt int

// Double is MyInt function
func (i MyInt) Double() int {
	// fmt.Printf("%T %v\n", i, i)
	// fmt.Printf("%T %v\n", 1, 1)
	// iはmain.MyIntなのでintでキャストしないといけない
	return int(i * 2)
}

// non structはあまりつかわない
func main() {
	myInt := MyInt(10)
	fmt.Println(myInt.Double())
}
