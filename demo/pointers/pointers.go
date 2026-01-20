package main

import "fmt"

type counter struct {
	hits int
}

func intrement(counter *counter) {
	counter.hits += 1
	fmt.Println("Counter", counter)
}

func replace(old *string, new string, counter *counter) {
	*old = new
	intrement(counter)

}

func main() {
	counter := counter{}

	hello := "Hello"
	world := "World"
	fmt.Println(hello, world)

	replace(&hello, "Hi", &counter)
	fmt.Println(hello, world)

	phrase := []string{hello, world}
	fmt.Println(phrase)

	replace(&phrase[1], "Go!", &counter)
	fmt.Println(phrase)
}
