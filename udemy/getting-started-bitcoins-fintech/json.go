package main

import (
	"encoding/json"
	"fmt"
)

type T struct{}

// json.Mashalの形を指定できる
type Person struct {
	// Age       int      `json:"age,string"` // ,型で型を指定できる
	Name      string   `json:"name,omitempty"` // name == ""なら表示しない
	Age       int      `json:"age,omitempty"`  // age == 0なら表示しない
	Nicknames []string `json:"nicknames,omitempty"`
	T         *T       `json:"T,omitempty"` // 独自の型はポインタ型にしないとomitemptyが効かない
}

// Marshalをカスタマイズできる
// func (p Person) MarshalJSON() ([]byte, error) {
// v, err := json.Marshal(&struct {
// 	Name string
// }{
// 	Name: "Mr." + p.Name,
// })
// return v, err
// }

// Unmarshalも同様にカスタマイズ
func (p *Person) UnmarshalJSON(b []byte) error {
	type Person2 struct {
		Name      string
		Age       int
		Nicknames []string
	}

	var p2 Person2
	err := json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println(err)
	}
	// ここでカスタマイズ
	p.Name = p2.Name + "============"
	p.Age = p2.Age + 100
	p.Nicknames = p2.Nicknames
	return err
}

func main() {
	// jsonをstructに入れるときのやり方
	// 大文字小文字どちらでも判定してくれるのがjson.Unmarshal
	b := []byte(`{"Name": "mike", "age":0, "nicknames": ["a", "b", "c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	// structをjsonに変更
	v, _ := json.Marshal(p)
	fmt.Println(string(v))
}
