package main

import "fmt"

func main() {
	channel := make(chan int, 2)

	channel <- 0
	channel <- 7

	close(channel)
	for range 3 {
		value, ok := <-channel
		fmt.Println("Получено:", value, "Ok:", ok)
	}
}
