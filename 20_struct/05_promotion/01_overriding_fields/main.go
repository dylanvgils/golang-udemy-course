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

func main() {
	p1 := hero{
		person: person{
			firstname: "John",
			lastname:  "Doe",
			age:       20,
		},
		firstname:     "Bond",
		licenseToKill: true,
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

	fmt.Println(p1.firstname, p1.person.firstname)
	fmt.Println(p2.firstname, p2.person.firstname)
}
