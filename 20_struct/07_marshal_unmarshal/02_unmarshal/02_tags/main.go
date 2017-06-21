package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Firstname string
	Lastname  string
	Age       int `json:"wisdome score"`
}

func main() {
	var p1 person

	bs := []byte(`{"firstname": "John", "lastname": "Doe", "wisdome score": 20}`)
	json.Unmarshal(bs, &p1)

	fmt.Println(p1)
}
