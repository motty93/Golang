package main

import "fmt"

func thirdParatyConnectDB() {
	// panicは非推奨
	panic("unable to connect database")
}

func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdParatyConnectDB()
}

func main() {
	save()
	fmt.Println("before do panic process")
}
