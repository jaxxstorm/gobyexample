package main

import "fmt"

//define a struct, person
type person struct {
	name string
	age  int
}

func main() {

	//Create a new struct
	fmt.Println(person{"Bob", 20})

	//Name the fields when creating new struct
	fmt.Println(person{name: "Alice", age: 30})

	//Omitted field is zero valued
	fmt.Println(person{name: "Fred"})

	//& will yield a pointer to a struct
	fmt.Println(&person{name: "Ann", age: 40})

	//assign value of struct to s
	s := person{name: "Sean", age: 50}

	//access struct fields with a dot
	fmt.Println(s.name)

	//structs with pointers. Need to learn more
	sp := &s
	fmt.Println(sp.age)

	//structs are mutable
	sp.age = 51
	fmt.Println(sp.age)

}
