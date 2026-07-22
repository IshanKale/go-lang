package main

import (
	"fmt"
	"time"
)
//channel is lilke await in javascript
func main() {
	// 1. Create a new channel that can transport string data
	messages := make(chan string,6)

	// 2. Start a concurrent goroutine using an anonymous function
	go func() {
		defer close(messages)
		for i:=0;i<5;i++{
			// Simulate some quick background work
			time.Sleep(500 * time.Millisecond)
	
			// Send a value into the channel using the <- operator
			messages <- "ping"
		}
	}()

	fmt.Println("Waiting for a message...")

	// 3. Receive the value from the channel (this blocks until data arrives)
	// msg := <-messages only one redived
	for msg := range messages {
		fmt.Println("Received:", msg)
	}

	// 4. Print the received message
}
