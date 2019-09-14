/*
 Channels are the pipes that connect concurrent goroutines.
 You can send values into channels from one goroutine and
 receive those values into another goroutine.
*/

package main

import "fmt"

func main() {
	// Create a new channel with make(chan val-type).
	// Channels are typed by the values they convey.
	messages := make(chan string)

	//Send a value into a channel using the channel <- syntax.
	go func() { messages <- "ping" }()

	//The <-channel syntax receives a value from the channel.
	msg := <-messages
	fmt.Println(msg)
}

// By default sends and receives block until
// both the sender and receiver are ready.
