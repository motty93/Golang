package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// structをPrintlnした際にageまでも表示したくない場合、String()インスタンスメソッドを実装すると隠せる
func (p Person) String() string {
	return fmt.Sprintf("my name is %v.", p.Name)
}

func main() {
	mike := Person{"mike", 20}
	fmt.Println(mike)
}
