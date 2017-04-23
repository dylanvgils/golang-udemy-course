package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 23
	fmt.Printf("%T - %v\n", myAge, myAge)
}

/*
	Here we declared a new type, foo
	the underlying type of foo is: int

	This is an EXAMPLE of a user defined type,
	it is BAD practice to alias types
		one exeption: if you jou need to attach methods to a type
*/
