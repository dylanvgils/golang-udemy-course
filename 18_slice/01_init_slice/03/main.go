package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 3)

	fmt.Println("--------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("--------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("Len: %d\t Capacity: %d\t Value: %d\n", len(mySlice), cap(mySlice), mySlice[i])
	}
}
