package main

import (
	"fmt"
	"time"
)

func funcA(chA chan<- string) {
	time.Sleep(1 * time.Second)
	chA <- "funcA finished"
}

func funcB(chB chan<- string) {
	time.Sleep(2 * time.Second)
	chB <- "funcB finished"
}

func main() {
	chA := make(chan string)
	chB := make(chan string)
	defer close(chA)
	defer close(chB)
	finflagA := false
	finflagB := false
	go funcA(chA)
	go funcB(chB)
	for {
		select {
		case msg := <-chA:
			finflagA = true
			fmt.Println(msg)
		case msg := <-chB:
			finflagB = true
			fmt.Println(msg)
		}

		if finflagA && finflagB {
			break
		}
	}
}
