package main

import "fmt"

func main () {

  // make a slice of strings of length 3 (zero valued)
  s := make([]string, 3)

  fmt.Println("emp:", 3) //print the empty slice

  s[0] = "a"
  s[1] = "b"
  s[2] = "c"
  fmt.Println("set:", s) //set the values of the slice
  fmt.Println("get:", s[2]) //get a value

  fmt.Println("len:", len(s)) //get the length, just like an array

  s = append(s, "d") // not available in arrays
  s = append(s, "e", "f") // returns a new slice with new values
  fmt.Println("apd:", s)

  c := make([]string, len(s)) // make a new slice, same length as s

  copy(c, s) // copy the values of s slice to c slice
  fmt.Println("cpy:", c)

  l := s[2:5] // slice operator. Get the elements from 2 - 5 in s slice
  fmt.Println("sl1:", l)

  l = s[:5] // up to, but excluding, s[5]
  fmt.Println("sl2:", l)

  l = s[2:] // up from and including s[2]
  fmt.Println("sl3:", l)

  t := []string{"g", "h", "i"} //decalare slice and initializa vars in single line
  fmt.Println("dcl:", t)

  twoD := make([][]int, 3) //two dimensional slice
  for i := 0; i < 3; i++ {
    innerLen := i + 1
    twoD[i] = make([]int, innerLen)
    for j := 0; j < innerLen; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2d: ", twoD)

}
