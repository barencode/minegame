package main

import "fmt"

func main() {
	channel := make(chan int)
	total := 0

	go worker(channel, 3)

	for received := range channel {
		total += received
		fmt.Println("Поступило:", received, "всего:", total)
	}
}

func worker(channel chan int, amount int) {
	for range 3 {
		channel <- amount
	}
	close(channel)
}
