package main

import "fmt"

// format here is: (param1 type, param2, type) return_type
func plus(a int, b int) int {

	return a + b

}

// all params have same type, so only use it once
func plusPlus(a, b, c int) int {

	return a + b + c

}

func main() {

	res := plus(1, 2)

	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

}
