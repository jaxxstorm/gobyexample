// this is go's only type of loop, cool!

package main // binary, of course

import "fmt"

func main () {

  i := 1 // becomes an integer

  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }

  fmt.Println("starting j loop")

  // start at 7, increase by 1 each iteration until we're at 9
  for j := 7; j <=9; j++ {
    fmt.Println(j)
  }

  for { 
    fmt.Println("loop")
    break
  }

}
