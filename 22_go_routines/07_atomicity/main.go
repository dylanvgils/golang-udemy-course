package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

func incrementor(s string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 50; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final counter:", counter)
}

/*
	go run -race main.go
	vs
	go run main.go
*/
