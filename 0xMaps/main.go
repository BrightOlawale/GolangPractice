// Description: Learning about Maps in go

package main

import "fmt"

func main() {
	fmt.Println("Let's learn about Maps")

	// Map is used to create something like a dictionary with key value pair
	languages := make(map[string]string) // if i want the value to be an integer, we use map[string][int]

	// Add key-value pair
	languages["En"] = "English"
	languages["Fn"] = "French"
	languages["Ch"] = "Chinese"
	languages["Ger"] = "German"
	languages["Esp"] = "Spanish"

	fmt.Println("All languages: ", languages)

	// Delete something from the map
	delete(languages, "Ch")
	fmt.Println("Edited languages: ", languages)

	// Loops thru the map
	for key, value := range languages {
		fmt.Printf("For key %v, the value is %v \n", key, value)
	}
}
