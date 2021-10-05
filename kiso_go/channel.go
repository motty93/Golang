package main

import "fmt"

func forCount(s chan<- int, num int) {
	for i := 0; i < num; i++ {
		s <- i
	}

	close(s)
}

func main() {
	c := make(chan int)
	go forCount(c, 5)

	for {
		value, ok := <-c

		if !ok {
			break
		}

		fmt.Println(value)
	}
}
