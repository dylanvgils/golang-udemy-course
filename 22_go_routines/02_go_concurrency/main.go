package main

import "fmt"

func foo() {
	for i := 0; i < 50; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 50; i++ {
		fmt.Println("Bar:", i)
	}
}

func main() {
	go foo()
	go bar()
}
