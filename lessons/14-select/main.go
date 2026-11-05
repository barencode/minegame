package main

import "fmt"

func main() {
	first := make(chan int, 1)
	second := make(chan int, 1)

	first <- 3
	second <- 5

	for range 2 {
		select {
		case value := <-first:
			fmt.Println("Из первого канала пришло:", value)
		case value := <-second:
			fmt.Println("Из второго канала пришло:", value)
		}
	}
}
