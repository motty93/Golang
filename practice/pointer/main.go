package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// &を使うとポインタ型を生成できる
	// &pはpのアドレスを表す
	p := Person{
		Name: "taro",
		Age:  20,
	}

	fmt.Printf("p :%+v\n", p)

	// 新しくp2というpと似たオブジェクトを生成している
	// 値渡し
	p2 := p
	p2.Name = "jiro"
	p2.Age = 19
	fmt.Printf("書き換えたはずのp :%+v\n", p)

	// &pで*Personを生成する
	// p3はpのアドレスが格納されている状態になる
	// 参照渡し, 構造体のフィールドの情報を変更するときにこれをよく使う
	p3 := &p
	p3.Name = "jiro"
	p3.Age = 19
	fmt.Printf("アドレスを書き換えたp :%+v\n", p)
}
