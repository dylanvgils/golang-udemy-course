package main

import "fmt"

func main() {
	x := 44
	changeMe(x)
	fmt.Println(x)
}

func changeMe(z int) {
	fmt.Println(z)
	z = 24
}

/*
  When chageMe is called on line 8
  The value 44 is being passed as an argument.
*/
