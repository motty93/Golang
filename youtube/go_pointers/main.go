package main

import "fmt"

type User struct {
	email    string
	username string
	age      int
	file     []byte
}

func getUser() (*User, error) {
	return nil, fmt.Errorf("not found")
}

func (u User) Email() string {
	return u.email
}

// *Userじゃないと更新されない→値渡しになってしまうので参照渡しにする必要あり
func (u *User) UpdateEmail(email string) {
	u.email = email
}

// pointer = 8byte
// x amount of bytes => sizeOf(user)となるので値渡しのメソッドだとメモリの無駄
func Email(u *User) string {
	return u.email
}

func main() {
	user := User{email: "test@foo.com", username: "test", age: 20}
	fmt.Println(user.Email())
	user.UpdateEmail("test2@foo.com")
	// NOTE: Email(user)はメモリが無駄
	fmt.Println(Email(&user))
}
