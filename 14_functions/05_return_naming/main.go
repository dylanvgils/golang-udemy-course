package main

import "fmt"

func main() {
	fmt.Println(greet("Jane", "Doe"))
}

func greet(fname, lname string) (s string) {
	s = fmt.Sprintf("%s %s", fname, lname)
	return
}
