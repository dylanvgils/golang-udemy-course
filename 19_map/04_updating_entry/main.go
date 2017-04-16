package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"John": "Good morning.",
		"Doe":  "Bonjour.",
	}

	myGreeting["Foo"] = "Howdy"
	fmt.Println(myGreeting)
	myGreeting["Foo"] = "Hi"
	fmt.Println(myGreeting)
}
