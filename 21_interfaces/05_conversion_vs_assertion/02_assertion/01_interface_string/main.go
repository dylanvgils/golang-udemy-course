package main

import "fmt"

func main() {
	// var name := "John" // Does not work (No interface)
	var name interface{} = "John"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Println("Value is not a string")
	}
}
