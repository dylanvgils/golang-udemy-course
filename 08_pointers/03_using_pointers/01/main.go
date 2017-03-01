package main

import "fmt"

func main() {
	a := 43
	fmt.Println(a)  // 43
	fmt.Println(&a) // Memory address

	var b = &a
	fmt.Println(b)  // 43
	fmt.Println(*b) // Memory address

	*b = 42
	fmt.Println(a) // 42
}

/*
  This is usefull,
  We can pass a memory address instead of a bunch of values (we can pass a reference)
  and then we can still change the value of whatever is stored at that memory address.
  This makes our programs more performant.
  We don't have to pass around large amounts of data,
  we only have to pass around adresses.

  everything is PASS BY VALUE in go.
  When we pass a memory address, we are passing a value.
*/
