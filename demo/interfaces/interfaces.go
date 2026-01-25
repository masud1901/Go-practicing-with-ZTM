package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("Preparing chicken dish:", c)
}

func (s Salad) PrepareDish() {
	fmt.Println("chop salad")
	fmt.Println("Add dishes")
	fmt.Println("Preparing salad dish:", s)
}

func prepareDishes(dishes []Preparer) {
	fmt.Println("Prepare dishes:")
	for i, dish := range dishes {
		fmt.Printf("Dish %d:\n", i+1)
		dish.PrepareDish()
	}

}

func main() {
	dishes := []Preparer{
		Chicken("Grilled Chicken"),
		Salad("Caesar Salad"),
		Chicken("Fried Chicken"),
		Salad("Greek Salad"),
	}

	prepareDishes(dishes)
}
