package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	age       int
}

type hero struct {
	person
	firstname     string
	licenseToKill bool
}

func (p person) greeting() {
	fmt.Println("I'm just a ragular person.")
}

func (h hero) greeting() {
	fmt.Println("I'am a super hero!")
}

func main() {
	p1 := person{
		firstname: "John",
		lastname:  "Doe",
		age:       20,
	}

	p2 := hero{
		person: person{
			firstname: "Jane",
			lastname:  "Doe",
			age:       24,
		},
		firstname:     "Jané",
		licenseToKill: false,
	}

	p1.greeting()
	p2.greeting()
	p2.person.greeting()
}
