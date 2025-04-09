//--Summary:
//  Create a program that can perform dice rolls. The program should
//  report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:

//* The program must handle any number of dice, rolls, and sides
//
//--Notes:
//* Use packages from the standard library to complete the project

package main

import (
	"fmt"
	"math/rand"
)

func roll(sides int) int {
	return rand.Intn(sides) + 1
}

func main() {

	dice, sides := 5, 6

	rolls := 2

	for r := 1; r <= rolls; r++ {
		sum := 0

		for d := 0; d < dice; d++ {
			rolled := roll(sides)
			sum += rolled
			fmt.Println("Roll #", r, ", Die #", d, ":", rolled)

		}
		fmt.Println("Total rolled:", sum)

		switch sum := sum; {
		case sum == 2 && dice == 2:
			//  - "Snake eyes": when the total roll is 2, and total dice is 2
			fmt.Println("Snake Eyes!")
			//  - "Lucky 7": when the total roll is 7
		case sum == 7:
			fmt.Println("Lucky 7")
			//  - "Even": when the total roll is
		case sum%2 == 0:
			fmt.Println("Even")
			//  - "Odd": when the total roll is odd
		case sum%2 == 1:
			fmt.Println("Odd")
		}
	}

}
