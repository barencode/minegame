package main

import "fmt"

func main() {
	count := 0
	fmt.Println("main: до вызова", count)

	go worker()

	count++
	fmt.Println("main: после вызова", count)
}

func worker() {
	fmt.Println("работаю")
}
