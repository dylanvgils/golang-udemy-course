package main

import "fmt"

func main() {
	greet := func() {
		fmt.Println("Hello world")
	}

	greet()
	fmt.Printf("Type: %T\n", greet)
}
