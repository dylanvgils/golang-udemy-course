package main

import "fmt"

func main() {
	x := 0

	// Func expression
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}

/*
  Anonymous function:
  function without a name

  Func expression:
  Assigning a function to a variable
*/
