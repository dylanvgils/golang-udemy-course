package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Get the book
	res, err := http.Get("http://www.gutenberg.org/files/2701/2701-0.txt")
	if err != nil {
		log.Fatalln(err)
	}
	// Scan the pages
	sc := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function for the scanning operation
	sc.Split(bufio.ScanWords)
	// Create slice to hold counts
	buckets := make([]int, 250)
	// Loop over words
	for sc.Scan() {
		n := hashBucket(sc.Text())
		buckets[n]++
	}

	fmt.Println(buckets[65:123])
}

func hashBucket(word string) int {
	return int(word[0])
}
