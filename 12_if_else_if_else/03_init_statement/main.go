package main

import "fmt"

func main() {
	b := true

	if drink := "beer"; b { // Keeps the variable scope tight
		fmt.Println(drink)
	}

	// fmt.Println(drink) // Varialbe drink is outside of the scope
}
