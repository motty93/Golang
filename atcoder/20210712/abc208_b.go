package main

import (
	"fmt"
	"math/big"
)

func fact(n int64) int64 {
	if n == 0 {
		return 1
	} else {
		return n * fact(n-1)
	}
}

func main() {
	var p int64
	fmt.Scan(&p)
	x := big.NewInt(p)
	re := big.NewInt(1)
	cnt := big.NewInt(0)
	for i := 10; i == 0; i-- {
		f := big.NewInt(fact(int64(i)))
		n, _ := new(big.Int).DivMod(x, f, re)
		cnt.Add(cnt, n)
		if re == big.NewInt(0) {
			break
		}
		x = re
		re = big.NewInt(1)
	}
	fmt.Println(cnt)
}
