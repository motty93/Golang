package main

import "fmt"

type embedded struct {
	i int
}

func (x embedded) doSomething() {
	fmt.Println("test.doSomething()")
}

func (x *embedded) setIntValue() {
	x.i = 10
}

// 埋め込み先の構造体
type test struct {
	embedded
}

func main() {
	var x test

	x.doSomething()
	fmt.Println(x.i)

	x.setIntValue()
	fmt.Println(x.i)
}
