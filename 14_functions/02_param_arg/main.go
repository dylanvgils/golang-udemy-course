package main

import "fmt"

func main() {
	greet("Jane")
	greet("John")
}

func greet(name string) {
	fmt.Println("Hello", name)
}

/*
  Greet is declared with a parameter
  When calling greet, pass in an argument
*/
