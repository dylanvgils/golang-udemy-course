package main

import "fmt"

// Variable x initialized in package scope
var x = 1

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}
