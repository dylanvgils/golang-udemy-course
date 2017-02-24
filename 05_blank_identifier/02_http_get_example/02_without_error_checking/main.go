package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("http://placehold.it")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s\n", page)
}

/*
  '_' is blank identifier in go, with the blank identifier you are telling
  the compiler throw away the returned value.
*/
