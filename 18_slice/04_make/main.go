package main

import "fmt"

func main() {
	customNumbers := make([]int, 3)
	// 3 is the length & capacity
	// length - number of elements referred to by the slice
	// capacity - number of alements in underlying array
	customNumbers[0] = 7
	customNumbers[1] = 15
	customNumbers[2] = 20

	fmt.Println(customNumbers[0])
	fmt.Println(customNumbers[1])
	fmt.Println(customNumbers[2])

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array
	// you could also do it like this

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "dias!"

	fmt.Println(greeting[2])
}
