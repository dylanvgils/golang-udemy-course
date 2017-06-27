package main

import "fmt"

type vehicles interface{}

type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct {
	vehicle
	Wheels int
	Doors  int
}

type plane struct {
	vehicle
	Jet bool
}

type boat struct {
	vehicle
	Length int
}

func main() {
	prius := car{}
	tesla := car{}
	bmw528 := car{}
	boeing747 := plane{}
	boeing777 := plane{}
	malibu := boat{}
	rides := []vehicles{prius, tesla, bmw528, boeing747, boeing777, malibu}

	for key, value := range rides {
		fmt.Println(key, "-", value)
	}
}
