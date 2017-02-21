package main

import "fmt"

// Wrapper has a function as it's return type
func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	// fmt.Println(x) // x is outside the scope for the main function
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
