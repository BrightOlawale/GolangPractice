// Description: Learning about Methods in Golang

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's look into Methods")

	newUser := User{"Olatnji", "tunji@dev.go", 18, true}
	fmt.Println("New User:", newUser)
	fmt.Printf("New User details: %+v\n", newUser)
	fmt.Printf("New User's name is %v\n", newUser.Name)

	// Execute the methods
	newUser.CheckName()
	newUser.ChangeMail("bright@go.dev")

	// Email doesn't change since the method created a COPY
	fmt.Printf("New User's Email is %v\n", newUser.Email)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

// Defining a method to print user's name: Note that anything not capitalized in the struct or method
// won't be exported.
func (user User) CheckName() {
	fmt.Println("The name is:", user.Name)
}

// A method to change the email: note that, it is going to be a copy since it;'s now pointing to the instance
func (u User) ChangeMail(newEmail string) {
	u.Email = newEmail
	fmt.Println("New Email is:", u.Email)
}
