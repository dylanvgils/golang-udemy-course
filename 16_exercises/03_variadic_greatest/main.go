package main

import "fmt"

func main() {
	greatest := max(10, 25, 33, 45, 100, 878, 1, 2)
	fmt.Println(greatest)
}

func max(numbers ...int) int {
	var largest int

	for _, v := range numbers {
		if largest < v {
			largest = v
		}
	}

	return largest
}
