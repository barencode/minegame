package main

import "fmt"

func main() {
	channel := make(chan int)
	close(channel)

	for range 2 {
		select {
		case value, ok := <-channel:
			fmt.Println(value, ok)
			if ok == false {
				channel = nil
			}
		default:
			fmt.Println("Нет готовых операций")
		}
	}
}
