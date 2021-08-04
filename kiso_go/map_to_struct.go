package main

import (
	"encoding/json"
	"fmt"
)

type Addr struct {
	PostalCode int
	Country    string
}

type User struct {
	Name    string
	Age     int
	Address Addr
	Email   string
}

func main() {
	user := map[string]interface{}{
		"Name": "motty",
		"Age":  30,
		"Address": map[string]interface{}{
			"PostalCode": 8113209,
			"Country":    "japan",
		},
		"Email": "test@gmail.com",
	}

	// mapをjsonへ変換
	byte, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byte))

	// jsonをstructへ割当
	var u User
	if err := json.Unmarshal(byte, &u); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user)
}
