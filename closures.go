package main

import "fmt"

// define an anonymous function
// essentially, this is a nameless function inside another function
// formation function name, parameters (a function) and then type
func intSeq() func() int {
	//define a var (int)
	i := 0
	// return another function
	return func() int {
		// plus one to i inside the function
		i += 1
		// then return i
		return i
	}
}

func main() {

	//Call the outer function
	nextInt := intSeq()

	//because i is one originally, and we call it multiple times
	//this basically increases i each call
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// start from scratch
	newInts := intSeq()
	fmt.Println(newInts())

}
