package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := person{"Dylan", 23}
	fmt.Println(&p.name) // Memory address (0xc420010360)
	changeMe(&p)
	fmt.Println(p)       // &{Dylan 23
	fmt.Println(&p.name) // Memory address (0xc420010360)
}

func changeMe(z *person) {
	fmt.Println(z)       // &{Dylan 23
	fmt.Println(&z.name) // Memory address (0xc420010360)
	z.name = "Peter"
	fmt.Println(z)       // &{Peter 23}
	fmt.Println(&z.name) // Memory address (0xc420010360)
}
