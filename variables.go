package main

import "fmt"

func main () {

  var a string = "initial" // when we declare a var, we declare its type as well
  fmt.Println(a)

  var b, c int = 1, 2 // an integer, and multiple vars
  fmt.Println(b, c)

  var d = true // boolean values
  fmt.Println(d)

  var e int // an integer again
  fmt.Println(e)

  f := "short" // shorthand, this infers the type (in this case a string)
  fmt.Println(f)


}
