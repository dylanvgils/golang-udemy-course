package main

import "fmt"

func main() {
	switch "Ja" {
	case "Jan":
		fmt.Println("Hello, Jan")
	case "Peter":
		fmt.Println("Hello, Peter")
	case "Piet":
	default:
		fmt.Println("No friends?")
	}
}

/*
  Switch statement has no default fallthrough.
  fallthrough is optional.
  -- you can specify fallthrough by explicitly stating it
  -- break isn't needed like in other languages
*/
