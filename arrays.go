package main

import "fmt"

func main() {

	var a [5]int           // an array, 5 items long, all nul valueds
	fmt.Println("emp:", a) // prints the while array? cool! no iteration needed!

	a[4] = 100 //set the value of the fourth element to 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a)) // length of the array

	b := [5]int{1, 2, 3, 4, 5} // give the elements default values. One liner.
	fmt.Println("dcl:", b)

	var twoD [2][3]int // a two dimensional array
	for i := 0; i < 2; i++ {
		//nested loops
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
