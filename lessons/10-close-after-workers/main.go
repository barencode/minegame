package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan int)
	var wg sync.WaitGroup
	total := 0

	wg.Go(func() {
		worker(channel, 3)
	})

	wg.Go(func() {
		worker(channel, 5)
	})

	go closeWhenDone(channel, &wg)

	for received := range channel {
		total += received
		fmt.Println("Поступило:", received, "всего:", total)
	}
}

func worker(channel chan int, amount int) {
	for range 3 {
		channel <- amount
	}
}

func closeWhenDone(channel chan int, wg *sync.WaitGroup) {
	wg.Wait()
	close(channel)
}
