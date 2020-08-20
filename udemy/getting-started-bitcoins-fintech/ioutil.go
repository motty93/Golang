package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// ioutilに関してはioに特化してる、特にネットワーク関係
func main() {
	// content, err := ioutil.ReadFile("main.go")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(content))
	//
	// if err := ioutil.WriteFile("iotuil_temp.go", content, 0666); err != nil {
	// 	log.Fatalln(err)
	// }

	r := bytes.NewBuffer([]byte("abc"))
	content, _ := ioutil.ReadAll(r)
	fmt.Println(string(content))
}
