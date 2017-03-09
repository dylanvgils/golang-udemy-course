package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

/*
  Exercise 5: Print all of the even numbers between 0 and 100.
*/
