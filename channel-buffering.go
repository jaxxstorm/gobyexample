package main

import "fmt"

// by default, channels are unbuffered, meaning they willl only accept sends (chan <-) if
// there is a corresponding receive (<- chan) ready.

// Buffered channels accept a limited number of values without any receiver being ready

func main() {

	// make a new channel
	// it can buffer up to 2 values
	messages := make(chan string, 2)

	// we know it's buffered, so we can pipe stuff into it
	// without a corresponding receive channel
	messages <- "buffered"
	messages <- "channel"

	// and the results are still there as we're expect
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
