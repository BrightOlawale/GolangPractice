// Description: Working with Web requests in Go

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Defining the url to be used
const url = "https://www.wikipedia.org/"

func main() {
	fmt.Println("Learning about Web request in Golang")

	// Making a GET request to the url using the net/http package
	response, err := http.Get(url)

	// Checking for errors
	if err != nil {
		panic(err)
	}

	// Printing the response type
	fmt.Printf("Response Type: %T\n", response)

	// Reading the response body into a byte slice
	dataByte, err := ioutil.ReadAll(response.Body)

	// Printing the response body in string format
	fmt.Println(string(dataByte))
}
