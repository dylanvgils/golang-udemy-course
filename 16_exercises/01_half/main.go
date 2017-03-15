package main

import "fmt"

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
}

func half(x int) (int, bool) {
	return x / 2, x%2 == 0
}
