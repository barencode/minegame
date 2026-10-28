package main

import "fmt"

func main() {
	channel := make(chan int)

	go worker(channel, 3)
	go worker(channel, 5)
	first := <-channel
	second := <-channel

	fmt.Println(first, second, "and total:", first+second)
}

func worker(result chan int, amount int) {
	result <- amount
}
