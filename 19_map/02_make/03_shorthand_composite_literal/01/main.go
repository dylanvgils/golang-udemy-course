package main

import "fmt"

func main() {
	myGreeting := map[string]string{}
	myGreeting["John"] = "Good morning."
	myGreeting["Doe"] = "Bonjour."

	fmt.Println(myGreeting)
}
