package main

import "fmt"

func main() {
	switch "Piet" {
	case "Jan":
		fmt.Println("Hello, Jan")
	case "Peter":
		fmt.Println("Hello, Peter")
	case "Piet":
		fmt.Println("Hello, Piet")
		fallthrough
	case "Co":
		fmt.Println("Hello, Co")
		fallthrough
	case "Chris":
		fmt.Println("Hello, Chris")
	default:
		fmt.Println("No friends?")
	}
}
