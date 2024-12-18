package main

import (
	"fmt"
	"os/user"
	"time"
)

var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func main() {
	fmt.Println("vim-go", time.Now())
	fmt.Println(user.Current())

	// 型推論は関数内でしか使えない
	xi := 1
	// xf64 := 1.2
	var xf64 float32 = 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("%T\n", xf64)
	fmt.Printf("%T\n", xi)
}
