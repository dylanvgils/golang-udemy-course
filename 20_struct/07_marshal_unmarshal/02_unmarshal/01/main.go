package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Firstname string
	Lastname  string
	Age       int
}

func main() {
	var p1 person
	fmt.Println(p1)

	bs := []byte(`{"firstname": "John", "lastname": "Doe", "age": 20}`)
	json.Unmarshal(bs, &p1)

	fmt.Println(p1)
	fmt.Printf("%T \n", p1)
}
