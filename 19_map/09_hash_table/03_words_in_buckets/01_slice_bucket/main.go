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
	buckets := make([][]string, 12)
	// Loop over words
	for sc.Scan() {
		word := sc.Text()
		n := hashBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}
	// Print length of each bucket
	for i := 0; i < 12; i++ {
		fmt.Printf("%d\t-\t%d\n", i, len(buckets[i]))
	}
}

func hashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
