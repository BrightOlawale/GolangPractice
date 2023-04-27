// Learning about Struct in golang

package main

import "fmt"

func main() {
	fmt.Println("Let's learn about struct")

	newUser := User{"Olatnji", "tunji@dev.go", 18, true}
	fmt.Println("New User:", newUser)
	fmt.Printf("New User details: %+v\n", newUser)
	fmt.Printf("New User's name is %v\n", newUser.Name)
}

// There is no inheritance in Golang or things like super
type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
