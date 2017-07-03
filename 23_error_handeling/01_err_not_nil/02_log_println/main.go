package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println(err)
	}
}

/*
	log.Println calls output to print to standard logger. Arguments are handeld in manner of fmt.Println
*/
