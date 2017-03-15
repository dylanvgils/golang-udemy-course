package main

import "fmt"

func main() {
	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["Dylan"])
}

func changeMe(z map[string]int) {
	z["Dylan"] = 23
}
