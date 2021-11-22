package main

import (
	"fmt"
	"os"
)

func main() {
	var sn, tn string
	fmt.Scan(&sn, &tn)
	if sn == tn {
		fmt.Println("Yes")
		os.Exit(0)
	}

	count := 0
	for i, s := range sn {
		if string(s) == tn[i:i+1] {
			count++
		}
	}

	if len(sn)-2 == count {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
