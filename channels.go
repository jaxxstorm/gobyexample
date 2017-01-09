package main

// channels allow us to connect goroutines

// we can essentially pipe things between goroutines

import "fmt"

func main() {

	// create a new channel with make
	// channels are typed by the values they convey
	messages := make(chan string)

	// send something into the channel using channel <- something
	// here we're sending the message ping into the messages channel
	go func() { messages <- "ping" }()

	// if you reverse the order, the channel receives the value
	// here we receive the ping message

	msg := <-messages
	fmt.Println(msg)

	// by default, send and receives will block until both are ready

}
