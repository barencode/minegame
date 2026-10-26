package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, result chan int, amount int, portions int, wg *sync.WaitGroup) {
	defer wg.Done()

	for range portions {
		if ctx.Err() != nil {
			fmt.Println("Работник: смена уже отменена")
			return
		}

		time.Sleep(time.Second)
		result <- amount
	}
}

func main() {
	demoCancel()

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	result := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)

	go worker(ctx, result, 1, 1, &wg)
	go worker(ctx, result, 5, 4, &wg)
	go closeWhenDone(result, &wg)

	coal := 0

	for received := range result {
		coal += received
		fmt.Println("Пришло:", received, "На смене:", coal)

		cancel()
	}
}

func closeWhenDone(result chan int, wg *sync.WaitGroup) {
	wg.Wait()
	close(result)
}

func demoCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	stop := ctx.Done()

	cancel()

	<-stop // GET: here op immediately stops
	fmt.Println("Сигнал отмены замечен")
}
