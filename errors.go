package main

import "errors"
import "fmt"

// go doesn't use exceptions, we define an explicit return value of type "error"

// by convention errors are the last return value and have type error
// this is a built in interface
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New constructs a basic error value with a message
		return -1, errors.New("Can't work with 42")
	}

	return arg + 3, nil // no error
}

// you can set custom types as errors by implement Error() on them

// define an argError struct
type argError struct {
	arg  int
	prob string
}

// here's our function returning a custom error
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// same as f1 essentially
func f2(arg int) (int, error) {
	if arg == 42 {

		// except we use the &argError struct for the error value
		return -1, &argError{arg, "can't work with this"}
	}
	return arg + 3, nil
}

func main() {

	// zero the return value
	// we only want the error
	for _, i := range []int{7, 42} {
		// this is common in go, check for no error
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)

	// replace standard error value with our custom error
	// using assertion
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}
