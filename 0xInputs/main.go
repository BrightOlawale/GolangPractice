// Description: Learning to accept user input to stdin so it could be used in my code

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	greeting := "Welcome to today's Class"

	fmt.Println(greeting)

	createdReader := bufio.NewReader(os.Stdin)
	fmt.Println("What was your score in the last test?")

	response, _ := createdReader.ReadString('\n') // This is comma ok || comma err syntax
	fmt.Println("You can always do better than ", response)
	fmt.Printf("And the type of this response is a %T", response)
}
