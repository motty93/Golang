package main

import "fmt"

func main() {

	i := 0
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			// rubyでいうnext
			continue
		}
		fmt.Println(n)
	}
}
