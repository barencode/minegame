package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	first := make(chan int)
	second := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	total := 0

	wg.Go(func() {
		worker(ctx, first, 3)
	})

	wg.Go(func() {
		worker(ctx, second, 5)
	})

	value := <-first
	fmt.Println("Получено в первом канале:", value)
	total += value

	value = <-second
	fmt.Println("Получено во втором канале:", value)
	total += value

	defer fmt.Println("Итого накоплено:", total)
	cancel()

	wg.Wait()
}

func worker(ctx context.Context, channel chan int, amount int) {
	defer close(channel)

	for range 3 {
		select {
		case channel <- amount:
			fmt.Println("Отправлено:", amount)
		case <-ctx.Done():
			fmt.Println("Работник остановлен:", amount)
			return
		}
	}
}
