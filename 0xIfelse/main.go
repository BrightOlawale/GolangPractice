// Description: Learning about If and Else statement in Golang

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("If and Else statement in Go")

	fmt.Println("Enter User Login count")
	createdReader := bufio.NewReader(os.Stdin)
	response, _ := createdReader.ReadString('\n')
	intResponse, _ := strconv.ParseFloat(strings.TrimSpace(response), 64)
	intVal, _ := strconv.ParseInt(strings.TrimSpace(response), 10, 64)

	if intResponse > 10 {
		fmt.Println("Regular user")
	} else if intResponse < 10 {
		fmt.Println("Not a regular")
	}

	if intVal%2 == 0 {
		fmt.Println("Login count is even")
	} else {
		fmt.Println("Login count is odd")
	}
}
