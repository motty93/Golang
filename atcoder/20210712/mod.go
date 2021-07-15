package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewInt(100)
	y := big.NewInt(30)
	z := big.NewInt(1)
	n, nn := new(big.Int).DivMod(x, y, z)
	fmt.Println(n)
	fmt.Println(nn)
	fmt.Println(z)
}
