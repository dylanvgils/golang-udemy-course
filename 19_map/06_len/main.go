package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"John": "Good morning.",
		"Doe":  "Bonjour.",
		"Foo":  "Howdy.",
	}

	fmt.Println(len(myGreeting))
}
