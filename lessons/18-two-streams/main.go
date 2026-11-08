package main

import "fmt"

func main() {
	first := make(chan int)
	second := make(chan int)

	total := 0

	go worker(first, 3)
	go worker(second, 5)

	for first != nil || second != nil {
		select {
		case value, ok := <-first:
			if !ok {
				first = nil
			} else {
				total += value
				fmt.Println("Получение из первого канала:", value, ok, "Всего накоплено:", total)
			}
		case value, ok := <-second:
			if !ok {
				second = nil
			} else {
				total += value
				fmt.Println("Получение из второго канала:", value, ok, "Всего накоплено:", total)
			}
		}
	}
}

func worker(channel chan int, amount int) {
	for range 3 {
		channel <- amount
	}
	close(channel)
}
