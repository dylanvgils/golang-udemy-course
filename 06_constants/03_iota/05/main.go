package main

import "fmt"

const (
	_  = iota             // Ignore first value by assigning blank identifier
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB                    // 1 << (2 * 10)
	GB                    // 1 << (3 * 10)
	TB                    // 1 << (4 * 10)
)

func main() {
	fmt.Println("Decimal\t\tBinary")
	fmt.Printf("%d\t\t%b\n", KB, KB)
	fmt.Printf("%d\t\t%b\n", MB, MB)
	fmt.Printf("%d\t%b\n", GB, GB)
	fmt.Printf("%d\t%b\n", TB, TB)
}
