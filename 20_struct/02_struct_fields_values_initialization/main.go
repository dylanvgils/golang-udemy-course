package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	age       int
}

func main() {
	p1 := person{"John", "Doe", 20}
	p2 := person{"Jane", "Doe", 24}
	fmt.Println(p1.firstname, p1.lastname, p1.age)
	fmt.Println(p2.firstname, p2.lastname, p2.age)
}
