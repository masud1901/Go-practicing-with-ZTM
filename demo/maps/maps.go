package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)

	shoppingList["eggs"] = 11
	shoppingList["milk"] = 34
	shoppingList["bread"] = 11
	shoppingList["bread"] += 11

	shoppingList["egg"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")

	fmt.Println("Milk deleted, new list:", shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]

	fmt.Println("Need cereal?")

	if !found {
		fmt.Println("nope")
	} else {
		fmt.Println("yup", cereal, "boxes")
	}

	totalItems := 0

	for _, amount := range shoppingList {
		totalItems += amount
	}

	fmt.Println("There are", totalItems, "items on the shopping list")
}
