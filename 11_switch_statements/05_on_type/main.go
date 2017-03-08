package main

import "fmt"

type contact struct {
	greeting string
	name     string
}

// SwitchOnType determines the type of provide value
func SwitchOnType(x interface{}) {
	switch x.(type) { // this is an assert; asserting, "x is of this type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	SwitchOnType(42)
	SwitchOnType("Dylan")

	var c = contact{"Hello", "Peter"}
	SwitchOnType(c)
	SwitchOnType(c.greeting)
	SwitchOnType(c.name)
}
