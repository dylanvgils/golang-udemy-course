package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)
	foo()
}

func foo() {
	// No access to x, so this does not compile
	fmt.Println(x)
}
