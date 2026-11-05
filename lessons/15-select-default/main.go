package main

import "fmt"

func main() {
	channel := make(chan int, 1)

	for range 2 {
		select {
		case value := <-channel:
			fmt.Println("Получено:", value)
		default:
			fmt.Println("Данных пока нет")
			channel <- 7
		}
	}
}
