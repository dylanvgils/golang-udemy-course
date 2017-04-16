package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"John": "Good morning.",
		"Doe":  "Bonjour.",
		"Foo":  "Howdy.",
	}

	fmt.Println(myGreeting)

	// delete(myGreeting, "Foo")

	if val, exists := myGreeting["Foo"]; exists {
		fmt.Println("That value exists")
		fmt.Println("val:", val)
		fmt.Println("exits:", exists)
	} else {
		fmt.Println("That value doesn't exist")
		fmt.Println("val:", val)
		fmt.Println("exits:", exists)
	}
}
