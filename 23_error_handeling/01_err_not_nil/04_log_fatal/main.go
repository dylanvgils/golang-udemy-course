package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

/*
	log.Fatalln is equivalent to Println() followed by a call to os.Exit(1).
*/
