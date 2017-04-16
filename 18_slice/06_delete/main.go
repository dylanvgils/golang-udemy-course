package main

import "fmt"

func main() {
	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wensday", "Thursday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)

	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println(mySlice)
}
