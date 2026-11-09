package main

import (
	"fmt"
	"sync"
)

func main() {
	first := make(chan int)
	second := make(chan int)
	stop := make(chan bool)
	var wg sync.WaitGroup
	total := 0

	wg.Go(func() {
		worker(first, stop, 3)
	})
	wg.Go(func() {
		worker(second, stop, 5)
	})

	value := <-first
	fmt.Println("Получено в первом канале:", value)
	total += value

	value = <-second
	fmt.Println("Получено во втором канале:", value)
	total += value

	defer fmt.Println("Итого накоплено:", total)
	close(stop)

	wg.Wait()
}

func worker(channel chan int, stop chan bool, amount int) {
	defer close(channel)

	for range 3 {
		select {
		case channel <- amount:
			fmt.Println("Отправлено:", amount)
		case <-stop:
			fmt.Println("Работник остановлен:", amount)
			return
		}
	}
}
