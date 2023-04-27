// Description: Learning about Switch cases in Golang

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in Golang")

	rand.Seed(time.Now().UnixMicro()) // sets the seed value for the random number generator

	dice := rand.Intn(7) // Selects a random number between 0-6 (7 not included)

	fmt.Println("Dice rolled, produced:", dice)

	switch dice {
	case 1:
		fmt.Println("Move one step forward")
	case 2:
		fmt.Println("Move two step forward")
	case 3:
		fmt.Println("Move three step forward")
	case 4:
		fmt.Println("Move four step forward")
		fallthrough // Goes to the next case after running this case
	case 5:
		fmt.Println("Move five step forward")
	case 6:
		fmt.Println("Move six step forward and roll the dice")
	default:
		fmt.Println("Do nothing my friend!")
	}
}
