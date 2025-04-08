package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 9, 3, 4

	if quiz1 > quiz2 {
		fmt.Println("Quiz1 scored higher than Quiz2")
	} else if quiz1 < quiz2 {
		fmt.Println("Quiz2 scored higher than Quiz1")

	} else {
		fmt.Println("Quiz 1 and quiz 2 have the same score")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("Acceped")
	} else {
		fmt.Println("Not Accepted")
	}

}
