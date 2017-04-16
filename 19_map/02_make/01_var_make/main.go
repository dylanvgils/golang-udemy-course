package main

import "fmt"

func main() {
	var myGreeting = make(map[string]string)
	myGreeting["John"] = "Good morning."
	myGreeting["Doe"] = "Bonjour."

	fmt.Println(myGreeting)
}
