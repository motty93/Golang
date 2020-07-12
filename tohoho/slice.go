package main

import "fmt"

func display(str string) {
	if str == "red" {
		fmt.Println("Reddddddddddddddddd")
	} else {
		fmt.Println(str)
	}
}

// 可変な配列のことをsliceと呼ぶ
func main() {
	a1 := []string{}
	a1 = append(a1, "red")
	a1 = append(a1, "green")
	a1 = append(a1, "black")
	a1 = append(a1, "blue")

	for i := 0; i < len(a1); i++ {
		display(a1[i])
	}
}
