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
	buckets := make([]int, 12)
	// Loop over words
	for sc.Scan() {
		n := hashBucket(sc.Text(), 12)
		buckets[n]++
	}

	fmt.Println(buckets)
}

func hashBucket(word string, buckets int) int {
	return int(word[0]) % buckets
}
