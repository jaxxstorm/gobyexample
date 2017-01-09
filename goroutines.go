package main

import "fmt"

// goroutines are lightweight execution threads
// they make it easy to make multi-threaded apps

// a simple function that prints the passed string
// 3 times
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// this is the normal way, run it once synchronously
	f("direct")

	// to execute it concurrently, do this
	go f("goroutine")

	// you can use goroutines on anonymous function calls
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// in order to check the output order
	// we require a keypress before exiting
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
