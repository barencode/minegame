package main

import "fmt"

func main() {
	channel := make(chan int)
	go worker(channel)
	value := <-channel
	fmt.Println(value)
}

func worker(result chan int) {
	result <- 7
}
