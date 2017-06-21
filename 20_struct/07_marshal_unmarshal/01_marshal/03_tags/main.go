package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Firstname string
	Lastname  string `json:"-"`
	Age       int    `json:"wisdome score"`
}

func main() {
	p1 := person{"John", "Doe", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}

/*
	Unsing "-" will ignore the field in the json object.
*/
