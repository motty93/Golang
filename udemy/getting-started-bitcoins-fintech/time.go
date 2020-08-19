package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	// postgresでよく使われるformat
	fmt.Println(t.Format(time.RFC3339))
}
