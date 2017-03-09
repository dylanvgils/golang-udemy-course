package main

import "fmt"

func main() {
	var largeNumber int
	var smallNumber int

	fmt.Print("Fill in a small number: ")
	fmt.Scan(&smallNumber)

	fmt.Print("Fill in a large number: ")
	fmt.Scan(&largeNumber)

	fmt.Printf("Deviding %d / %d remainder is %d\n", largeNumber, smallNumber, largeNumber%smallNumber)
}

/*
  Exercise 4: Create a program that prints to the terminal asking for a user to
  enter a small number and a larger number. Print the remainder of the larger
  number divided by the smaller number.
*/
