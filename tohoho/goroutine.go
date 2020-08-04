package main

import (
	"fmt"
	"time"
)

func routineTest(num int) {
	for i := 0; i < 10; i++ {
		fmt.Print("A")
		time.Sleep(time.Duration(num) * time.Millisecond)
	}
}

func main() {
	go routineTest(10)
	for i := 0; i < 10; i++ {
		fmt.Print("M")
		time.Sleep(time.Duration(20) * time.Millisecond)
	}
}
