package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	age       int
}

func (p person) fullName() string {
	return fmt.Sprintf("%s %s", p.firstname, p.lastname)
}

func main() {
	p1 := person{"John", "Doe", 20}
	p2 := person{"Jane", "Doe", 24}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
