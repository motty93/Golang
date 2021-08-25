package main

import "fmt"

func f1() (int, string, float32) {
	return 0, "xy", 3.14

}

func f2(a int, b string, c interface{}) {
	fmt.Println(a, b, c)
}

func main() {
	// returnした多値をそのままfunctionへ渡す
	f2(f1())
}
