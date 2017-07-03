package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println(err)
	}
}

/*
	fmt.Println fomrts using the default formats and writes to standard output.
*/
