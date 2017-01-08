package main

import "fmt"
import "math"

// interfaces are named sets of methods

// declare an interface
//
type geometry interface {
	area() float64
	perim() float64
}

// a struct for rectangles
type rect struct {
	width, height float64
}

// and circles
type circle struct {
	radius float64
}

// to implement an interface, create an instance of all methods in it
// so first, let's implement the area method

func (r rect) area() float64 {
	return r.width * r.height
}

// and now the perim method
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// note, we havem't had to declare the individual methods here, nice!

// now for circles

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// if you declare a variable of type "interface"
// you can use any of the methods in the interface

func measure(g geometry) {
	fmt.Println(g)         // returns the struct values
	fmt.Println(g.area())  // returns the area
	fmt.Println(g.perim()) // returns the perimeter
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
