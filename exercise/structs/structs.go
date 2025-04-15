//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//

//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,

//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter

type coordinate struct {
	x, y int
}
type rectangle struct {
	topleft     coordinate
	bottomright coordinate
}

func width(rect rectangle) int {
	return (rect.bottomright.x - rect.topleft.x)

}

func length(rect rectangle) int {
	return (rect.topleft.y - rect.bottomright.y)
}

func area(rect rectangle) int {
	return length(rect) * width(rect)
}

func perimeter(rect rectangle) int {
	return 2 * (length(rect) + width(rect))
}

func printInfo(rect rectangle) {
	fmt.Println("Area is", area(rect))
	fmt.Println("Perimeter is", perimeter(rect))
}

func main() {

	rect := rectangle{topleft: coordinate{0, 7}, bottomright: coordinate{10, 0}}

	printInfo(rect)
	//* After performing the above requirements, double the size
	//  of the existing rectangle and repeat the calculations
	//  - Print the new results to the terminal

	rect.topleft.y *= 2
	rect.bottomright.x *= 2

	printInfo(rect)

}
