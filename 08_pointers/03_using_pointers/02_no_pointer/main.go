package main

import "fmt"

func zero(z int) {
	fmt.Printf("%p\n", &z) // address in func zero
	z = 0
}

func main() {
	x := 5
	fmt.Println(&x) // address in main
	zero(x)
	fmt.Println(x) // x is still 5
}
