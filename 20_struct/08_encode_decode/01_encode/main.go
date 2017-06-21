package main

import (
	"encoding/json"
	"os"
)

type person struct {
	Firstname string
	Lastname  string
	Age       int
}

func main() {
	p1 := person{"John", "Doe", 20}
	json.NewEncoder(os.Stdout).Encode(p1)
}
