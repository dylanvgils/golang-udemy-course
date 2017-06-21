package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Firstname   string
	Lastname    string
	Age         int
	notExported int
}

func main() {
	p1 := person{"John", "Doe", 20, 007}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}
