package main

import "fmt"

func main() {

  m := make(map[string]int) // make a map, zero values (map[key-type]val-type])

  m["k1"] = 7
  m["k1"] = 13

  fmt.Println("map:", m) // printing maps is easy!

  v1 := m["k1"] // make v1 the value of the map

  fmt.Println("v1: ", v1) // print it

  fmt.Println("len:", len(m)) // length of the map

  delete(m, "k2") // delete something from a map
  fmt.Println("map:", m)

  _, prs := m["k2"] // blank identifier. I don't need the value, so I have a second return value. This is seems to be essentially a nice if value.exist? from ruby
  fmt.Println("prs:", prs) // will return false, because the value doesn't exist

  n := map[string]int{"foo": 1, "bar": 2}
  fmt.Println("map:", n)
                 


}
