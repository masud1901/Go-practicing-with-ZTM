//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:

//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import (
	// "flag"
	"fmt"
)

//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter



 type coordinates struct {
	x, y float64
 }
type Rectangle struct {
	a coordinates
	b coordinates
}

func area(a coordinates, b coordinates) float64 {
	length := b.x - a.x
	width := b.y - a.y
	return length * width
}

func perimeter(a coordinates, b coordinates) float64 {
	length := b.x - a.x
	width := b.y - a.y
	return 2 * (length + width)
}

func main() {
	rect := Rectangle{
		a: coordinates{x: 0, y: 0},
		b: coordinates{x: 4, y: 3},
	}

	areaResult := area(rect.a, rect.b)
	perimeterResult := perimeter(rect.a, rect.b)
	fmt.Printf("Area: %.2f\n", areaResult)
	fmt.Printf("Perimeter: %.2f\n", perimeterResult)

	// Double the size of the rectangle
	rect.b.x *= 2
	rect.b.y *= 2

	areaResult = area(rect.a, rect.b)
	perimeterResult = perimeter(rect.a, rect.b)
	fmt.Printf("New Area: %.2f\n", areaResult)
	fmt.Printf("New Perimeter: %.2f\n", perimeterResult)

}