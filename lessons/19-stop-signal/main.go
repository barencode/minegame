package main

import (
	"fmt"
	"sync"
)

func main() {
	result := make(chan int)
	stop := make(chan int)
	var wg sync.WaitGroup

	wg.Go(func() {
		worker(result, stop, 3)
	})

	value := <-result
	fmt.Println("Получено:", value)
	close(stop)

	wg.Wait()
}

func worker(result chan int, stop chan int, amount int) {
	defer close(result)

	for range 3 {
		select {
		case result <- amount:
			fmt.Println("Отправлено:", amount)
		case <-stop:
			fmt.Println("Работник остановлен")
			return
		}
	}
}
