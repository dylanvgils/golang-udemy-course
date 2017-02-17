package main

import (
	"fmt"

	"bitbucket.org/dylanvgils/udemy-golang-course/04_scope/01_package_scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.Firstname)
	// Will not work because lastname is not exported
	//fmt.Println(vis.lastname)
	vis.PrintVars()
}
