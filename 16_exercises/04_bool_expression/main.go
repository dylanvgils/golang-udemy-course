package main

import "fmt"

func main() {
	result := (true && false) || (false && true) || !(true && false) // true
	fmt.Println(result)
}
