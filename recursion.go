package main

import "fmt"

//define a function fact with param n(int) and returns an int
func fact(n int) int {

	// let's not multiply by 0
	if n == 0 {
		return 1
	}
	//recursive function!
	//we're multiplying the function BY ITSELF
	return n * fact(n-1)
}

func main() {

	fmt.Println(fact(7))

}
