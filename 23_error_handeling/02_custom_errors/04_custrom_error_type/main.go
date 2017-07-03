package main

import (
	"fmt"
	"log"
)

type SqrtMathError struct {
	str string
	err error
}

func (n *SqrtMathError) Error() string {
	return fmt.Sprintf("Error occured: (%v) %v", n.str, n.err)
}

func main() {
	_, err := sqrt(-1)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(n int64) (int64, error) {
	if n < 0 {
		msg := fmt.Errorf("Square root of negative number: %v", n)
		return 0, &SqrtMathError{"Hello world", msg}
	}
	return 42, nil
}
