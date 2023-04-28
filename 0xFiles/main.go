// Learning how to work with files in Golang

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Let's look into files in Golang")

	// Creating what to write to file
	joke := ` Welcome to my Quantum Physics Joke
	Heisenberg and Schrödinger get pulled over for speeding.
	The cop asks Heisenberg "Do you know how fast you were going?"
	Heisenberg replies, "No, but we know exactly where we are!"
	The officer looks at him confused and says "you were going 108 miles per hour!"
	Heisenberg throws his arms up and cries, "Great! Now we're lost!"
	The officer looks over the car and asks Schrödinger if the two men have anything in the trunk.
	"A cat," Schrödinger replies.
	The cop opens the trunk and yells "Hey! This cat is dead."
	Schrödinger angrily replies, "Well he is now."`

	// Creating the file to be written into. Something like "touch physics_jokes.txt"
	file, err := os.Create("./physics_jokes.txt")

	// If err terminate program
	checkForError(err)

	// Write the created joke string into the physics_jokes.txt file
	length, err := io.WriteString(file, joke)

	checkForError(err)

	// Print the length of the file to screen
	fmt.Println("Length of the file:", length)

	// Close the file, I can declare this line after creating the file since I used the defer keyword
	defer file.Close()

	// Reading the file to screen
	readFile("./physics_jokes.txt")
}

// Let's Read a that file and print to screen
func readFile(file string) {
	fmt.Println("--------------------------------------------------------------------------------")

	// This ioutil.ReadFile method will return a Byte not the original readable text
	dataInByte, err := ioutil.ReadFile(file)

	checkForError(err)

	// Let us convert the data from Byte to Readable string
	dataInString := string(dataInByte)

	fmt.Println("Joke #1:", dataInString)
}

// Creating a check error function
func checkForError(err error) {
	// Terminate program if error occurs
	if err != nil {
		panic(err)
	}
}
