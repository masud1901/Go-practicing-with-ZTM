//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:

package main

import (
	"fmt"
	"math/rand"
)

//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func displayName(name string) {
	fmt.Println("Welcome", name)
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()

func message() string {
	return "Hello, Hows it going?"
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func addNumbers(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

//* Write a function that returns any number

func returnAnyNumber() int {
	num := rand.Int()
	return num
}

//* Write a function that returns any two numbers

func returnAnyTwoNumber() (int, int) {
	return rand.Int(), rand.Int()
}

func main() {
	name := "Ayon"
	displayName(name)

	fmt.Println(message())
	num1, num2 := returnAnyTwoNumber()
	num3 := returnAnyNumber()

	//* Add three numbers together using any combination of the existing functions.
	//  * Print the result

	answer := addNumbers(num3, num1, num2)
	fmt.Println("The three random numbers are", num1, num2, num3)

	fmt.Println("The three random numbers adition", answer)

	//* Call every function at least once

}
