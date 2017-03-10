package main

import "fmt"

func main() {
	fmt.Println(greet("Jane", "Doe"))
}

func greet(fname, lname string) (string, string) {
	return fmt.Sprintf("%s %s", fname, lname), fmt.Sprintf("%s %s", lname, fname)
}
