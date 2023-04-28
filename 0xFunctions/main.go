// Desription: Learning about Functions in Golang

package main

import "fmt"

func main() {
	fmt.Println("A lesson on Function")

	// NOTE: You do not define a function inside another function

	greetings("Tunji")

	value := adder(5, 8)
	fmt.Println("Answer is:", value)

	totalResult := totalAdder(8, 2, 4, 9, 6, 1, 5)
	fmt.Println("Total result is:", totalResult)

	age, email := twoReturn(18, "Coco@dev.go")
	fmt.Println("Age:", age, "Email:", email)
}

// Function to print A Greeting to screen
func greetings(name string) {
	fmt.Println("Good day", name)
}

// Function to add 2 values together
func adder(valA int, valB int) int { // The inr before the curly braces is used to tell the return type
	return valA + valB
}

// Function to accept multiple arguments and work on them
func totalAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

// Function to return to Types
func twoReturn(age int, email string) (int, string) {
	return age, email
}
