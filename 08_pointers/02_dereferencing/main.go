package main

import "fmt"

func main() {
	a := 43
	fmt.Println(a)  // 43
	fmt.Println(&a) // Memory address

	var b = &a
	fmt.Println(b)  // 43
	fmt.Println(*b) // Memory address
}

/*
  b is an int pointer;
  b points to the memory address where an int is stored.
  To see the value in that memory address, add a '*' in front of b,
  this is known as dereferencing.
  The '*' is an operator in this case.
*/
