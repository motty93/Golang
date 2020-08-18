package main

import "fmt"

type UserNotFound struct {
	Username string
}

// error実装の際はpointerを渡す
// error内容が同一かどうかわからないのでポインターにする
func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.Username)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "mike"}
}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}
