package main

import "fmt"

func main() {
	// Declaring variables without assignment
	var a string
	var b float64
	var c int
	var d bool

	fmt.Printf("%T - %v\n", a, a)
	fmt.Printf("%T - %v\n", b, b)
	fmt.Printf("%T - %v\n", c, c)
	fmt.Printf("%T - %v\n", d, d)
}

/*
  NOTE: By default variables are initialized to their default values,
  when no value is assigned. In Go we call this 'zero values'.

  String formatting (https://golang.org/pkg/fmt/)
  '%T' prints the type of the variable
  '%v' prints the actual value of the variable (Go figures out the type)
*/
