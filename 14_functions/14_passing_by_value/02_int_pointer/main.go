package main

import "fmt"

func main() {
	x := 44
	fmt.Println(&x) // Memory address (0xc42000a2f8)

	changeMe(&x)

	fmt.Println(&x) // Memory address (0xc42000a2f8)
	fmt.Println(x)  // 24
}

func changeMe(z *int) {
	fmt.Println(z)  // Memory address (0xc42000a2f8)
	fmt.Println(*z) // 44
	*z = 24
	fmt.Println(z)  // Memory address (0xc42000a2f8)
	fmt.Println(*z) // 24
}
