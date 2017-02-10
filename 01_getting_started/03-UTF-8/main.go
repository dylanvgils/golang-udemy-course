package main

import "fmt"

func main() {
	for i := 60; i < 122; i++ {
		fmt.Printf("Decimal: %d\tBinary: %b\tHexadecimal: %#X\t UTF-8: %q\n", i, i, i, i)
	}
}
