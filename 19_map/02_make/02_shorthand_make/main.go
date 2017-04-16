package main

import "fmt"

func main() {
	myGreeting := make(map[string]string)
	myGreeting["John"] = "Good morning."
	myGreeting["Doe"] = "Bonjour."

	fmt.Println(myGreeting)
}
