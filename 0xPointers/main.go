// Description: A lesson on Pointers in Golang

package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on Pointers")

	var num *int
	fmt.Println("Value of num = ", num)

	newNumner := 50 // Here is a variable holding the number 50

	var ptr = &newNumner // Here, the & is used to reference the address of the variable created above

	fmt.Println("The value of variable is: ", newNumner)
	fmt.Println("The address of the variable is: ", ptr)
	fmt.Println("The value of variable is: ", *ptr) // Here, I dereferenced that address to obtain
	// the value in the address

	// I can operate on the deferenced addressand it will affect the actual variable
	*ptr = *ptr * 2

	fmt.Println("The new value of the variable is: ", newNumner)
}
