package main

import "fmt"

func main() {
	channel := make(chan int)
	total := 0

	go worker(channel, 3)
	go worker(channel, 5)

	for range 6 {
		received := <-channel
		total += received
		fmt.Println("Поступило:", received, "всего:", total)
	}
}

func worker(channel chan int, amount int) {
	for range 3 {
		channel <- amount
	}
}
