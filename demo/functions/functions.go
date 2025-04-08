package main

import "fmt"

func double(number int) int {
	return 2 * number
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello from the greet Function!")
}

func main() {
	greet()
	num := 12214
	answer := double(num)
	fmt.Println("The Current number is", num, "and after the funnction the number is", answer)
	answer = add(answer, num)
	fmt.Println(answer)
	answer = add(double(answer), add(answer, num))
	fmt.Println(answer)

}
