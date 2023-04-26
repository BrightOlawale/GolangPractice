// Description: Learning about Time and Date in Golang

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Golang time class")

	currentTime := time.Now()

	fmt.Println(currentTime)

	fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday")) // This particular format i the only way to
	// get the current date in your desired Format

	createdDate := time.Date(2022, time.April, 22, 18, 52, 0, 0, time.UTC)

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("Monday 01-02-2006 15:04:05")
}
