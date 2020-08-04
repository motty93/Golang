package main

import "fmt"

type any interface{}

// pythonのdictもどきを定義できる
type dict map[string]any

func main() {
	p1 := dict{
		"name": "motty",
		"age":  90,
		"address": dict{
			"zip": "801-0112",
			"tel": "090-1111-2222",
		},
	}

	name := p1["name"]
	fmt.Println(name)
	// anyをdictに変換してから参照
	zip := p1["address"].(dict)["zip"]
	fmt.Println(zip)
	tel := p1["address"].(dict)["tel"]
	fmt.Println(tel)
}
