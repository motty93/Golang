package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// ポインタ方の変数を宣言, pがポインタ変数・*Personポインタ型
	var p *Person

	p = &Person{
		Name: "taro",
		Age:  20,
	}
	fmt.Printf("pのアドレス: %p", p)
	fmt.Printf("pのName: %v", p.Name)
}
