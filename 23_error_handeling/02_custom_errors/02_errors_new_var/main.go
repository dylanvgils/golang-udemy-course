package main

import (
	"errors"
	"log"
)

var ErrSqrt = errors.New("Square root of negative number")

func main() {
	//fmt.Printf("%T\n", errSqrt)
	_, err := sqrt(-1)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(n int64) (int64, error) {
	if n < 0 {
		return 0, ErrSqrt
	}
	return 42, nil
}
