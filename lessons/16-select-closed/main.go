package main

import "fmt"

func main() {
	channel := make(chan int)
	close(channel)

	for range 2 {
		select {
		case value, ok := <-channel:
			fmt.Println(value, ok)
		default:
			fmt.Println("Данных пока нет")
		}
	}
}
