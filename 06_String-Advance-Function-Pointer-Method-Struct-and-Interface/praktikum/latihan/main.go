package main

import "fmt"

type User struct {
	username  string
	email     string
	full_name string
}

// value receiver if i don't need to change the attribute value
func (u User) DisplayName() {
	fmt.Println(u.username)
}

// pointer receiver if i need to change the attribute 
func (u *User) SetUsername(newUsername string) {
	u.username = newUsername
}

func main() {
	person1 := User{"nibras", "nibras@example.com", "nibras alyassar"}

	person1.DisplayName()
	
	person1.SetUsername("sarbin")
	person1.DisplayName()
}