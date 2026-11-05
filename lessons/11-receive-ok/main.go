package main

import "fmt"

func main() {
	var channel = make(chan int)

	go worker(channel, 0, 7)

	for range 3 {
		value, ok := <-channel
		fmt.Println(value, ok)
	}

}

func worker(channel chan int, first int, second int) {
	channel <- first
	channel <- second
	close(channel)
}
