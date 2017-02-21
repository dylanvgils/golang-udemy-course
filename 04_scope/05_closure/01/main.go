package main

import "fmt"

// 'func main' is an outer scope for the inner scope
func main() {
	x := 42
	fmt.Println(x)

	// Inner scope
	{
		y := "Hello from the inner scope!"
		fmt.Println(y)
	}

	// fmt.Println(y) // Outside scope of 'y'
}

/*
	Closure helps us limit the scope of variables used by multiple functions
	without closure, for two or more functions to have access to the same variable,
	that variable would be needed to be package scope.
*/
