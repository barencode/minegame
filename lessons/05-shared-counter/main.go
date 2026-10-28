package main

import (
	"fmt"
	"sync"
)

var count int
var mu sync.Mutex

func main() {
	var wg sync.WaitGroup

	wg.Go(worker)
	wg.Go(worker)

	wg.Wait()

	fmt.Println(count)
}

func worker() {
	for range 100_000 {
		mu.Lock()
		count++
		mu.Unlock()
	}
}
