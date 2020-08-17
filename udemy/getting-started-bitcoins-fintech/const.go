package main

import "fmt"

const Pi = 3.14

const (
	Username = "test"
	Password = "test_pass"
)

// const だとoverflowしにくい
// var big int = 12340189730498127309 + 1
const big = 1234018973049812739 + 1

func main() {
	fmt.Println(Pi, Username, Password)
	fmt.Println(big - 1)
}
