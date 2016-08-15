package main

import "fmt"

// returns 2 values, both ints
func vals() (int, int) {
	return 3, 7
}

func main() {

	// a and b are equal to return values from vals func
	a, b := vals()

	fmt.Println(a)
	fmt.Println(b)

	// we only want a subset of the returned values
	_, c := vals()
	fmt.Println(c)

}
