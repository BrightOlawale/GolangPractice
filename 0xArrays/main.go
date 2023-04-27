// Description: Learning about Arrays in Golang

package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on Arrays in Golang")

	var myGuys [5]string

	myGuys[0] = "Victor"
	myGuys[2] = "Winner"
	myGuys[3] = "Samuel"

	fmt.Println("My guys are: ", myGuys)
	// Output: My guys are:  [Victor  Winner Samuel ]. Noticed the spaces in the missing index locations?
	fmt.Println("The length of the array is: ", len(myGuys))

	var numList = [4]int{1, 4, 7, 5}
	fmt.Println("Array has numbers: ", numList)
	fmt.Println("The length of the array is: ", len(numList))
}
