package main

import "fmt"

const greeting string = "Hello, %s!\n"

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Printf(greeting, name)
}
