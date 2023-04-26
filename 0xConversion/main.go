// Description: Learning how to convert strings to an Integer or Float that could be operated upon

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	greeting := "Hello, welcome to class"
	fmt.Println(greeting)

	fmt.Println("Please provide your score in the last test")

	createdReader := bufio.NewReader(os.Stdin)

	response, _ := createdReader.ReadString('\n')

	fmt.Println("Below is your Actual score;")

	actualScore, err := strconv.ParseFloat(strings.TrimSpace(response), 64)
	// If we didnt use the function string.TrimSpace on the response it could have only parsed
	// "4\n" to the ParseFloat function and printed the err in line 28

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(((actualScore / 10) * 100))
	}
}
