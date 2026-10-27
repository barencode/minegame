package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	count := 0
	fmt.Println("main: до вызова", count)

	wg.Go(worker)

	wg.Wait()
	fmt.Println("продолжаю работу")

	count++
	fmt.Println("main: после вызова", count)
}

func worker() {
	fmt.Println("работаю")
}
