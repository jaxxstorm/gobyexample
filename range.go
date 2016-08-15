// iterate over some shit
package main

import "fmt"

func main() {

  // create a slice
  nums := []int{2,3,4}

  //new integer, starting at zero
  sum := 0

  // sum the numbers in a slice
  // we don't need the index
  // so we blank it
  for _, num := range nums {
    sum += num
  }

  fmt.Println("sum:", sum)

  // here we need the index, so let's use it
  // ranges produce an index and a value
  for i, num := range nums {
    if num == 3 {
      fmt.Println("index:", i)
    }
  }

  kvs := map[string]string{"a": "apple", "b": "banana"} // new map
  // loop through the keys using range
  // assign to k, v
  for k, v := range kvs {
    fmt.Println("%s -> %s\n", k, v)
  }

  for i, c := range "go" {
    fmt.Println(i, c) // some crazy unicode witchcraft
  }

}
