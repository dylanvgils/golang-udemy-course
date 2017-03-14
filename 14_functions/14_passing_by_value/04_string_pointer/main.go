package main

import "fmt"

func main() {
	name := "Dylan"
	fmt.Println(&name) // Memory address (0xc42000a320)

	changeMe(&name)

	fmt.Println(&name) // Memory address (0xc42000a320)
	fmt.Println(name)  // Peter
}

func changeMe(z *string) {
	fmt.Println(z)  // Memory address (0xc42000a320)
	fmt.Println(*z) // Dylan
	*z = "Peter"
	fmt.Println(z)  // Memory address (0xc42000a320)
	fmt.Println(*z) // Peter
}
