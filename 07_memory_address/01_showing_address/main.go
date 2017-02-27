package main

import "fmt"

func main() {
	a := "Hello world!"

	fmt.Println("Value of a:", a)
	fmt.Println("Adderess of a:", &a)
}

/*
  With '&' you are able to get memory address of the variable.
*/
