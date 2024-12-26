package main

import "fmt"

func main() {
	myC := make(chan string)
	anotherC := make(chan string)

	go func() {
		myC <- "hello"
	}()

	go func() {
		anotherC <- "world"
	}()

	select {
	case msgFromMyChannel := <-myC:
		fmt.Println(msgFromMyChannel)
	case msgFromAnotherChannel := <-anotherC:
		fmt.Println(msgFromAnotherChannel)
	}
}
