package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Covert strings to value:
	a, _ := strconv.ParseBool("true")
	b, _ := strconv.ParseFloat("3.123", 64)
	c, _ := strconv.ParseInt("-42", 10, 64)
	d, _ := strconv.ParseUint("42", 10, 64)

	fmt.Println(a, b, c, d)

	// Convert values to strings
	w := strconv.FormatBool(true)
	x := strconv.FormatFloat(3.123, 'E', -1, 64)
	y := strconv.FormatInt(-42, 16)
	z := strconv.FormatUint(42, 16)

	fmt.Println(w, x, y, z)
}
