package main

import "fmt"
import "time"

// define a funciton with a param done
// this param is a channel which can notify
// another goroutine that the function is finished
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	// start a goroutine and specify a channel to notify
	done := make(chan bool, 1)
	go worker(done)

	// block/ don't do anything until we're done
	// we'll get a notification from the channel
	// if we removed this, the program would exit before
	// the worker ever started
	<-done
}
