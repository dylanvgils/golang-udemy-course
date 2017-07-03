package main

import (
	"errors"
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
		return 0, errors.New("Square root of negative number")
	}
	return 42, nil
}
