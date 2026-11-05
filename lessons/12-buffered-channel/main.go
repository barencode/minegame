package main

import "fmt"

func main() {
	var channel = make(chan int, 1)
	channel <- 7
	result := <-channel
	fmt.Println("Отправлено", result)

	channel <- 9
	result = <-channel
	fmt.Println("Отправлено", result)
}
