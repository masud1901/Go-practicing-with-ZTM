//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import (
	"fmt"
)

func anyNumber() int {
	return -1
}

func main() {
	var child string
	//* Use a `switch` statement to print the following:
	switch age := anyNumber(); {
	case age < 1:
		child = "newborn"
	case age < 4:
		child = "toddler"
	case age < 13:
		child = "child"
	case age < 18:
		child = "teenager"
	default:
		child = "adult"
	}
	fmt.Println("The guy based on the age is", child)

	//  - "newborn" when age is 0
	//  - "toddler" when age is 1, 2, or 3
	//  - "child" when age is 4 through 12
	//  - "teenager" when age is 13 through 17
	//  - "adult" when age is 18+

}
