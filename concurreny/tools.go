package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	//go helloWorld()
	fmt.Println("Sending to the channel")
	go func(input chan string) {
		input <- "Hello"
	}(c)
	fmt.Println("Receiving from the channel")
	greeting := <-c
	fmt.Printf("%v", greeting)

	time.Sleep(1 * time.Millisecond)
}

func helloWorld() {
	fmt.Println("Hello world")
}
