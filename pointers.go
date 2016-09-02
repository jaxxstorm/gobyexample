package main

import "fmt"

// takes a param ival which is an int
func zeroval(ival int) {
	ival = 0
}

//takes a POINTER to an int
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {

	// Initial value
	i := 1
	fmt.Println("initial:", i)

	//We basically print the int
	zeroval(i)
	fmt.Println("zeroval:", i)

	//Here's the good stuff
	//so above, the *iptr is i
	//by calling this function with the memory address of i
	//we dereference i (because it's the param *iptr)
	//which then changes the value
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	//print the pointer (memory address)
	fmt.Println("pointer:", &i)

}
