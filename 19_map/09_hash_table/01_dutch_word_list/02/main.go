package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://raw.githubusercontent.com/titoBouzout/Dictionaries/master/Dutch.dic")
	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[string]string)

	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words[sc.Text()] = ""
	}

	i := 0
	for k := range words {
		fmt.Println(k)

		if i >= 500 {
			break
		}

		i++
	}
}
