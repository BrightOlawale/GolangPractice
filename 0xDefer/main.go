// Description: Learning about the defer keyword in Golang

package main

import "fmt"

func main() {
	fmt.Println("Let's look into the defer keyword")

	// Anywhere the defer statement is called, it is kinda added into a LIFO queue
	// That is the first defer statement is executed last while the last is executed first,
	// The ones without the defer keyword are executed immediately
	defer fmt.Println("First line")
	defer fmt.Println("Second line")
	fmt.Println("Third line")
	defer fmt.Println("Fourth line")
	fmt.Println("Fifth line")
	defer fmt.Println("Last line")
	reverseString()
}

// Let's define a function and make it run in the main program
func reverseString() {
	characters := 5

	for characters > 0 {
		fmt.Print(characters, "\n")
		characters--
	}
}
