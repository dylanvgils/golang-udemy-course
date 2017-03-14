package main

import "fmt"

func main() {
	name := "Dylan"
	fmt.Println(name) // Dylan

	changeMe(name)

	fmt.Println(name) // Dylan
}

func changeMe(z string) {
	fmt.Println(z) // Dylan
	z = "Peter"
	fmt.Println(z) // Peter
}
