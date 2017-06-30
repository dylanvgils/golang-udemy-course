package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

/*
* Why does the code below only print zero?
* And what can you do to get it to print all 0 - 9 numbers?
 */

// func main() {
// 	c := make(chan int)

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			c <- i
// 		}
// 	}()

// 	fmt.Println(<-c)
// }
