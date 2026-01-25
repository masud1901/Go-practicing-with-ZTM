//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure\

//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

func checkRunes(line string) {
	letterCount := 0
	digitCount := 0
	spaceCount := 0
	punctCount := 0

	countRunes := func(r rune) {
		if unicode.IsLetter(r) {
			letterCount++
		} else if unicode.IsDigit(r) {
			digitCount++
		} else if unicode.IsSpace(r) {
			spaceCount++
		} else if unicode.IsPunct(r) {
			punctCount++
		}
	}
	for _, r := range line {
		countRunes(r)
	}

	fmt.Printf("Letters: %d\n", letterCount)
	fmt.Printf("Digits: %d\n", digitCount)
	fmt.Printf("Spaces: %d\n", spaceCount)
	fmt.Printf("Punctuation Marks: %d\n", punctCount)
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	for _, line := range lines {
		fmt.Println("Analyzing line:", line)
		checkRunes(line)
	}

}
