package main

import "fmt"

// takes an arbitray number of vars
func sum(nums ...int) {

	//print the numbers
	fmt.Print(nums, " ")
	total := 0
	//loop through the elements of the slice
	//add the them to the total
	for _, num := range nums {
		total += num
	}
	//Finally, print the total
	fmt.Println(total)
}

func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

}
