package main

import "fmt"

// define a struct
type rect struct {
	width, height int
}

// methods are function that operate on structs
// for laymen like me, you pass the struct as a param

// declare a method
// here we're using a pointer to *rect as the arg
func (r *rect) area() int {
	//both of these are ints, so you'll get an int back
	return r.width * r.height
}

// you can also use a value
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {

	// initialize the struct first
	r := rect{width: 10, height: 5}

	// operate on the defined r struct
	// looks a bit like OOP, but it's not
	fmt.Println("Area: ", r.area())

	fmt.Println("Perimeter: ", r.perim())

	// Go handles the value/pointer conversion when using methods
	// if you use a pointer, the method can manipulate the receiving struct
	rp := &r
	fmt.Println("Area: ", rp.area())
	fmt.Println("Perimeter: ", rp.perim())

}
