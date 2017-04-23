package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	age       int
}

type hero struct {
	person
	licenseToKill bool
}

func main() {
	p1 := hero{
		person: person{
			firstname: "John",
			lastname:  "Doe",
			age:       20,
		},
		licenseToKill: true,
	}

	p2 := hero{
		person: person{
			firstname: "Jane",
			lastname:  "Doe",
			age:       24,
		},
		licenseToKill: false,
	}

	fmt.Println(p1.firstname, p1.firstname, p1.age, p1.licenseToKill)
	fmt.Println(p2.firstname, p2.firstname, p2.age, p2.licenseToKill)
}
