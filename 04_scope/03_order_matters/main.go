package main

import "fmt"

func main() {
	// This will not work because in block level scope order matters
	fmt.Println(x)
	x := 10

	// This will work order in package level scope doesn't matter
	fmt.Println(y)
}

var y = 100
