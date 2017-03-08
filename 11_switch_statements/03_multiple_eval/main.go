package main

import "fmt"

func main() {
	switch "Jan" {
	case "Tim", "Jan":
		fmt.Println("Hello Tim or Jan")
	case "Chris", "Peter":
		fmt.Println("Hello Chris or Peter")
	}
}
