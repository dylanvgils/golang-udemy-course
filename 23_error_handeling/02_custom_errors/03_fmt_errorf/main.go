package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-1)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(n int64) (int64, error) {
	if n < 0 {
		return 0, fmt.Errorf("Square root of negative number: %v", n)
	}
	return 42, nil
}
