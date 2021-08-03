package main

import "fmt"

func main() {
	m := map[interface{}]interface{}{
		1:      10,
		"test": 20,
	}

	fmt.Println(m)
}
