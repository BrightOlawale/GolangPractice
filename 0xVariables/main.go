// Description: Learning about Variables and Types in Go

package main

import "fmt"

const UserEmail string = "oolawalebright@gmail.com" // Public

func main() {
	var username string = "oolatunji"
	fmt.Println(username)
	fmt.Printf("The above variable is of Type: %T\n", username)

	fmt.Println("**********************************************************************")

	var isRegistered bool = true
	fmt.Println(isRegistered)
	fmt.Printf("The above variable is of Type: %T\n", isRegistered)

	fmt.Println("**********************************************************************")

	var smallValue uint8 = 255 //(Unsigned 8-bit integers (0-255))
	fmt.Println(smallValue)
	fmt.Printf("The above variable is of Type: %T\n", smallValue)

	fmt.Println("**********************************************************************")

	var smallFloat float32 = 255.67812635584033075258
	fmt.Println(smallFloat)
	fmt.Printf("The above variable is of Type: %T\n", smallFloat)

	fmt.Println("**********************************************************************")

	// Default values and aliases
	var variableInt int
	fmt.Println(variableInt)
	fmt.Printf("The above variable is of Type: %T\n", variableInt)

	fmt.Println("**********************************************************************")

	// Implicit Type
	var name = "Olatunji"
	fmt.Println(name) // N.B: Since you passed a string you cannot just change the variable to a number,
	//only strings must be passed to the name variable.

	fmt.Println("**********************************************************************")

	// It is possible to ignore the var keyword
	totalCount := 5000.34688 // You are allowed to use ":=" inside any method/function but not outside it
	fmt.Println(totalCount)

	fmt.Println("**********************************************************************")

	fmt.Println(UserEmail)
	fmt.Printf("The above variable is of Type: %T\n", UserEmail)
}
