package main

import (
	"fmt"

	"github.com/dylanvgils/udemy-golang-course/02_package/stringutil"
)

func main() {
	// Not working because 'myName' is not exported in stringutil package.
	// fmt.Println(stringutil.myName)

	// Getting first and lastname through exported variables and methods.
	fmt.Printf("%s %s\n", stringutil.GetName(), stringutil.MyLastname)

	// Get the hello world string in the right order.
	fmt.Println(stringutil.Reverse("!dlrow olleH"))
}
