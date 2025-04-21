//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:

package main

import "fmt"

type Product struct {
	name  string
	price int
}

// * Print to the terminal:
func printProductInfo(product []Product, length int) {
	itemCount := len(product)

	//  - The last item on the list
	fmt.Println("The last item on the list", product[len(product)-1].name)
	//  - The total number of items
	fmt.Println("The total number of itmes", itemCount)
	//  - The total cost of the items
	totalCost := 0
	for i := 0; i < len(product); i++ {
		totalCost += product[i].price
	}
	fmt.Print("The total cost is ", totalCost)
	fmt.Println()
}

func main() {

	//* Using an array, create a shopping list with enough room
	//  for 4 products
	//  - Products must include the price and the name
	var shoppingList = [4]Product{}
	products := [...]Product{
		{name: "Rice", price: 45},
		{name: "Toy", price: 345},
		{name: "Barry", price: 5},
		{name: "Juice", price: 500},
	}
	//* Insert 3 products into the array
	for i := 0; i < 3; i++ {
		shoppingList[i] = products[i]
	}
	printProductInfo(shoppingList[:3], 3)

	//* Add a fourth product to the list and print out the
	//  information again
	shoppingList[3] = products[3]
	printProductInfo(shoppingList[:], len(shoppingList))
}
