package main

import "fmt"

func main() {
	var myName = "Masud"
	fmt.Println("Hello,my name is", myName)

	var name string = "janowar"
	fmt.Println("Name =", name)

	userName := "admin"
	fmt.Println("username =", userName)

	var sum int
	fmt.Println("The sum is", sum)

	part1, other := 1, 5
	fmt.Println("part 1 is", part1, "part 2 is", other)

	part2, other := 2, 0
	fmt.Println("part 1 is", part2, "part 2 is", other)

	sum = part1 + part2
	fmt.Println("sum =", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)

	fmt.Println("Lesson Name =", lessonName, "Lesson Type =", lessonType)

	word1, word2, _ := "Hello", "world", "!"

	fmt.Println(word1, word2)
}
