package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type person struct {
	Firstname string
	Lastname  string
	Age       int
}

func main() {
	var p1 person
	rdr := strings.NewReader(`{"firstname": "John", "lastname": "Doe", "age": 20}`)
	json.NewDecoder(rdr).Decode(&p1)
	fmt.Println(p1)
}
