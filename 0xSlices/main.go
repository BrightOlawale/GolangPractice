// Description: Learning about slices in go and what it us all about

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Let's learn about slices")

	var friendList = []string{} // Created slice here

	fmt.Printf("This variable 'friendList', is of type %T\n", friendList)

	friendList = append(friendList, "Winner", "Samuel", "Victor") // Added value to slice
	fmt.Println(friendList)                                       // output: [Winner Samuel Victor]

	friendList = append(friendList[0:2])
	fmt.Println(friendList) // Output: [Winner Samuel]

	classScore := make([]int, 3) // Create memory for slice using make

	classScore[0] = 299
	classScore[1] = 160
	classScore[2] = 1230
	// classScore[3] = 5000 // To return Index error as we only defined the slice with lenght of 3

	// But with append I can reallocate more memory to accomodate new values
	classScore = append(classScore, 200, 7000, 923)

	fmt.Println(classScore)
	fmt.Println("Slice is sorted?", sort.IntsAreSorted(classScore))

	// We can sort our slice using the methods below
	sort.Ints(classScore)

	fmt.Println(classScore)
	fmt.Println("Slice is sorted?", sort.IntsAreSorted(classScore))

	// Remove value from a slice
	var courses = []string{"CPE316", "CSC308", "CPE314", "CSC310", "CSC302"}
	fmt.Println("Courses offered: ", courses)
	indexToRemove := 3
	courses = append(courses[:indexToRemove], courses[indexToRemove+1:]...)
	fmt.Println("Courses now offered: ", courses)
}
