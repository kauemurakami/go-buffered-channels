package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	channel <- "Hello world"
	channel <- "Olá mundo"
	// channel <- "GO lang" //erro deadlock

	message := <-channel
	message2 := <-channel
	fmt.Println(message)
	fmt.Println(message2)

}
