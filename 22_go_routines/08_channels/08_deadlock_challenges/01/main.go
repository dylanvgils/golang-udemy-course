package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}

/*
* The code below results in a deadlock.
* Can you determine why?
* And what would you do to fix it?
*
 */

// func main() {
// 	c := make(chan int)
// 	c <- 1
// 	fmt.Println(<-c)
// }
