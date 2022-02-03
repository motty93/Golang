package main

import "fmt"

func primeFactorization(n int) (pf []int) {
	for n%2 == 0 {
		pf = append(pf, 2)
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pf = append(pf, i)
			n = n / i
		}
	}

	if n > 2 {
		pf = append(pf, n)
	}

	return
}

func main() {
	var n int
	fmt.Scan(&n)
	pfs := primeFactorization(n)
	fmt.Println(pfs)
	if len(pfs) > 2 || len(pfs) == 0 {
		fmt.Println("No")
	} else if len(pfs) == 1 && pfs[0] < 10 {
		fmt.Println("Yes")
	} else {
		fmt.Println("Yes")
	}
}
