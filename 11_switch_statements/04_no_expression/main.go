package main

import "fmt"

func main() {
	name := "Jan"

	switch {
	case len(name) == 2:
		fmt.Println("Hello name with lenght of 2")
	case name == "Jan":
		fmt.Println("Hello, Jan")
	case name == "Peter":
		fmt.Println("Hello, Peter")
	}
}
