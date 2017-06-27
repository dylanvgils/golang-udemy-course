package main

import "fmt"

func main() {
	rem := 3.14
	fmt.Printf("%T\n", rem)
	fmt.Printf("%T\n", int(rem))

	var val interface{} = 7
	fmt.Printf("%T\n", val)
	// fmt.Printf("%T\n", int(val)) // Does not work (Need type assertion)
	// fmt.Printf("%T\n", val.(int))
}
