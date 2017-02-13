package main

import "fmt"

// Declare
var a string

// Declare + Assign = Initialize
var b = "van Gils"

func main() {
	// Assign
	a = "Dylan"
	// Shorthand variable declaration
	c := "I'm a ninja"

	fmt.Printf("%T - %v\n", a, a)
	fmt.Printf("%T - %v\n", b, b)
	fmt.Printf("%T - %v\n", c, c)
}

/*
  NOTE: In Go it's not requird to specify the variable type,
  when not specified Go can figure it out.

	String formatting (https://golang.org/pkg/fmt/)
  '%T' prints the type of the variable
  '%v' prints the actual value of the variable (Go figures out the type)
*/
