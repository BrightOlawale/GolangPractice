// Description: Learning loops in Golang

package main

import "fmt"

func main() {
	fmt.Println("Loops in Go")

	// WE ONLY HAVE FOR LOOP IN GOLANG!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

	food := []string{"Amala", "Eba", "Semo", "Tuwo", "Pounded yam"}

	for n := 0; n < len(food); n++ {
		fmt.Println(food[n])
	}

	fmt.Println("****************************************************************************")
	// OR use range keyword to get the index  of elements in the slices
	for x := range food {
		fmt.Println(food[x])
	}

	fmt.Println("****************************************************************************")
	// OR get the index and the element using the for loop like below
	for idx, elem := range food {
		fmt.Printf("The index is %v and the value is %v\n", idx, elem)
	}

	fmt.Println("****************************************************************************")
	// Using continue and break statement
	count := 10
	for count > 0 {

		if count == 7 {
			goto end
		}

		fmt.Println(count)
		count--

		if count == 5 {
			count--
			continue
		}
	}

	// We can call end a label
end:
	fmt.Println("Jumped to the end.")
}
