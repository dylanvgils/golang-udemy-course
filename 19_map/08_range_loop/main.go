package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"John": "Good morning.",
		"Doe":  "Bonjour.",
		"Foo":  "Howdy.",
	}

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}
